package main

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"io"
	"os"
	"strings"
	"time"

	"github.com/czQery/PiFi/backend/api"
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func init() {

	var format = "02/01/2006 - 15:04:05"

	logrus.SetOutput(io.Discard)

	// Log terminal
	logrus.AddHook(&hp.LogFormatterHook{
		Writer: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:     true,
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: format,
		},
	})

	// Log file
	var err error
	hp.LogFile, err = os.OpenFile(hp.LogFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("init - log file load failed")
	}
	err = hp.LogFile.Truncate(0)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("init - log file truncate failed")
	}
	logrus.AddHook(&hp.LogFormatterHook{
		Writer: hp.LogFile,
		Formatter: &logrus.TextFormatter{
			ForceColors:   false,
			DisableColors: true,

			QuoteEmptyFields: true,
			ForceQuote:       true,
			DisableQuote:     false,

			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		},
	})
}

func main() {

	defer hp.LogFile.Close()

	hp.ArtPrint()
	hp.ConfigLoad()
	logrus.Info("config - successfully loaded")
	hp.DistLoad()

	nmInit()

	go ticker()

	r := fiber.New(fiber.Config{
		CaseSensitive:         false,
		DisableStartupMessage: true,
		GETOnly:               false,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ServerHeader:          "PiFi",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var e *api.Error
			if errors.As(err, &e) {
				log := logrus.WithFields(e.Fields())
				switch e.Code {
				case 401, 404, 503:
					break
				case 500:
					log.Error("fiber - " + e.Func)
				default:
					log.Warn("fiber - " + e.Func)
				}

				if e.Message != "" {
					return c.Status(e.Code).JSON(api.Response{Message: e.Message})
				}
			}

			return c.Status(500).JSON(api.Response{Message: "internal server error"})
		},
	})

	// API
	rAPI := r.Group("/api", func(c *fiber.Ctx) error {
		if !api.VerifyToken(c) {
			return &api.Error{Code: 401, Func: "api/settings", Message: "unauthorized"}
		}
		return c.Next()
	})

	rAPI.Get("/auth", api.Auth)
	rAPI.Get("/stats", api.Stats)
	rAPI.Get("/log", api.Log)
	rAPI.Get("/settings", api.SettingsGet)
	rAPI.Post("/settings", api.SettingsPost)

	// Pifi UI
	rUI := r.Group("/pifi", func(c *fiber.Ctx) error {

		if strings.TrimSuffix(c.Path(), "/") == "/pifi" {
			return c.Redirect("/pifi/dash", 308)
		}

		return c.Next()
	})

	rUI.All("/:tab", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/index.html")
	})
	rUI.All("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendFile("./dist/favicon.ico")
	})
	rUI.Static("/assets", "./dist/assets")
	rUI.Static("/font", "./dist/font")

	// Captive portal
	r.All("/", func(c *fiber.Ctx) error {
		return c.SendFile("./portal/test/index.html")
	})

	// Default
	r.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(string(c.Request().URI().Path()), "/api") {
			return &api.Error{Code: 404, Func: "api", Message: "unknown endpoint"}
		} else if !hp.Dist {
			return &api.Error{Code: 503, Func: "static", Message: "front-end unavailable"}
		} else {
			return c.Redirect("/", 307)
		}
	})

	logrus.Info("fiber - started")

	// Run
	err := r.Listen(hp.Config.String("main.address"))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("fiber - server failed")
	}
}

func nmInit() {
	// Get interfaces
	iface, ifaceErr := cmd.GetInterfaceList()
	if ifaceErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceErr.Error(),
		}).Panic("main - nmcli failed")
	}

	// Get interfaces from config
	ifaceConfig := make(map[string]map[string]interface{})
	ifaceConfigErr := hp.Config.Unmarshal("settings.iface", &ifaceConfig)
	if ifaceConfigErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceConfigErr.Error(),
		}).Panic("main - iface config load failed")
	}
	for _, item := range ifaceConfig {
		item["ready"] = false
	}

	// Create hotspot con for the first time
	initErr := errors.New("no wifi interface")
	for _, i := range iface {
		if i.Type == "wifi" {
			initErr = cmd.InitHotspot(i.Name)

			if item, e := ifaceConfig[i.Name]; e {
				item["ready"] = true
			} else {
				newItem := make(map[string]interface{})
				newItem["ready"] = true
				newItem["mode"] = "none"

				ifaceConfig[i.Name] = newItem
			}
		}
	}
	if initErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": initErr.Error(),
		}).Panic("main - hotspot init failed")
	}

	// Save edited interfaces config
	ifaceConfigErr = hp.Config.Set("settings.iface", ifaceConfig)
	if ifaceConfigErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceConfigErr.Error(),
		}).Panic("main - iface config save failed")
	}

	// Apply saved settings
	var settings api.SettingsResponse
	settingsErr := mapstructure.Decode(hp.Config.Get("settings"), &settings)
	if settingsErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": settingsErr.Error(),
		}).Panic("main - settings load failed")
	}

	applyErr := api.ApplySettings(settings)
	if applyErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": applyErr.Error(),
		}).Panic("main - settings apply failed")
	}

	logrus.Info("main - nmcli loaded")
}
