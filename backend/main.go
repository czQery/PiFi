package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
	"time"

	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/sse"
	"github.com/gofiber/fiber/v2/middleware/proxy"

	"github.com/mitchellh/mapstructure"

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

	if hp.Build == "dev" {
		logrus.SetLevel(logrus.DebugLevel)
	}

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
		Writer:    hp.LogFile,
		Broadcast: sse.DashLog,
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
	logrus.Info("main - Created by Štěpán Aubrecht")
	logrus.Info("main - Build: " + hp.Build)
	hp.ConfigLoad()
	logrus.Info("config - successfully loaded")
	hp.DistLoad()
	db.Load()

	go cmd.InitGPS()
	go cmd.InitBettercap()
	time.Sleep(time.Second * 5)

	nmInit()
	go Timer()

	r := fiber.New(fiber.Config{
		CaseSensitive:         false,
		DisableStartupMessage: true,
		GETOnly:               false,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ServerHeader:          cmd.Con,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var e *api.Error
			if errors.As(err, &e) {
				log := logrus.WithFields(e.Fields())
				switch e.Code {
				case 400, 401, 404, 503:
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
		if c.Path() != "/api/portal" && !api.VerifyToken(c) { // require password for all api endpoints except /api/portal
			return &api.Error{Code: 401, Func: "api", Message: "unauthorized"}
		}
		return c.Next()
	})

	rAPI.Get("/auth", api.Auth)
	rAPI.Get("/settings", api.SettingsGet)
	rAPI.Post("/settings", api.SettingsPost)
	rAPI.All("/portal", api.Portal)
	rAPI.Get("/portals", api.Portals)
	rAPI.Get("/db", api.DB)

	rAPI.Patch("/net/:net<regex(\\b(wigle|beacondb|dwpa)\\b)>", api.NetPatch)
	rAPI.Post("/net/:net<regex(\\b(wigle|beacondb|dwpa)\\b)>", api.NetPost)

	rAPI.Get("/net/wigle", api.WigleGet)
	rAPI.Get("/net/dwpa", api.DWPAGet)

	rAPI.Get("/attack/deauth", api.DeauthGet)
	rAPI.Post("/attack/deauth", api.DeauthPost)
	rAPI.Delete("/attack/deauth", api.DeauthDelete)

	rAPI.Get("/cap/:file", func(c *fiber.Ctx) error {
		return c.SendFile("./cap/" + c.Params("file"))
	})

	// Bettercap api proxy
	rAPI.All("/bettercap/events", proxy.Forward(cmd.BC+"/api/events"))
	rAPI.All("/bettercap/session", proxy.Forward(cmd.BC+"/api/session"))
	rAPI.Get("/bettercap/session/wifi", proxy.Forward(cmd.BC+"/api/session/wifi"))

	// SSE
	rSSE := r.Group("/sse", func(c *fiber.Ctx) error {
		if !api.VerifyToken(c) {
			return &api.Error{Code: 401, Func: "sse", Message: "unauthorized"}
		}
		return c.Next()
	})

	rSSE.Get("/dash", sse.Dash)

	// PiFi UI
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
	rPortal := r.Group("/", func(c *fiber.Ctx) error {

		c.Set("Cache-Control", "no-store")

		if c.Path() == "/" {
			if cmd.HotspotPortal == "" {
				return c.Redirect("/pifi/dash", 307)
			}
			return c.SendFile("./portal/" + cmd.HotspotPortal + "/index.html")
		}

		return c.Next()
	})
	rPortal.All("favicon.ico", func(c *fiber.Ctx) error {
		return c.SendFile("./portal/" + cmd.HotspotPortal + "/favicon.ico")
	})
	rPortal.All(":dir/:file", func(c *fiber.Ctx) error {
		return c.SendFile("./portal/" + cmd.HotspotPortal + "/" + c.Params("dir") + "/" + c.Params("file"))
	})

	// Default
	r.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(string(c.Request().URI().Path()), "/api") {
			return &api.Error{Code: 404, Func: "api", Message: "unknown endpoint"}
		} else if !hp.Dist {
			return &api.Error{Code: 503, Func: "static", Message: "front-end unavailable"}
		}

		return c.Redirect("/", 307)
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
	var (
		initHotspot         bool
		initHotspotSettings map[string]any
		initHotspotErr      error
	)

	// Get interfaces
	iface, ifaceErr := cmd.GetInterfaceList()
	if ifaceErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceErr.Error(),
		}).Panic("main - nmcli failed")
	}

	// Get interfaces from config & set them all as not ready before verifying if they are
	ifaceConfig := hp.ConfigGetInterfaceList()
	for _, item := range ifaceConfig {
		item["ready"] = false
	}

	connections, connectionErr := cmd.GetConnectionList()
	if connectionErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": connectionErr.Error(),
		}).Error("main - connections failed")
	}

	for _, con := range connections {
		if con.Name == cmd.Con+"-hotspot" {
			initHotspot = true
		}
	}

	for _, i := range iface {
		if i.Type != "wifi" {
			continue
		}

		if item, e := ifaceConfig[i.Name]; e {
			item["ready"] = true
			if item["mode"] == "hotspot" {
				initHotspot = true
			}
		} else {
			newItem := make(map[string]any)
			newItem["ready"] = true
			newItem["mode"] = "none"

			ifaceConfig[i.Name] = newItem
		}

		logrus.WithFields(logrus.Fields{
			"iface": i.Name,
			"state": i.State,
		}).Info("main - iface init")

		// Create hotspot con for the first time
		if !initHotspot && i.State == "disconnected" {
			ifaceConfig[i.Name], initHotspotSettings, initHotspotErr = cmd.InitHotspot(i.Name, ifaceConfig[i.Name])
			if initHotspotErr == nil {

				// Save new hotspot config
				initHotspotErr = hp.Config.Set("settings.hotspot", initHotspotSettings)
				if initHotspotErr != nil {
					logrus.WithFields(logrus.Fields{
						"err": initHotspotErr.Error(),
					}).Panic("main - hotspot config save failed")
				}

				initHotspot = true
			}
		}
	}

	if !initHotspot {
		logrus.WithFields(logrus.Fields{
			"err": "no wifi interface",
		}).Panic("main - hotspot init failed")
	}

	if initHotspotErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": initHotspotErr.Error(),
		}).Panic("main - hotspot init failed")
	}

	// Save edited interfaces config
	ifaceConfigErr := hp.Config.Set("settings.iface", ifaceConfig)
	if ifaceConfigErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceConfigErr.Error(),
		}).Panic("main - iface config save failed")
	}

	hp.ConfigSave()
	hp.ConfigLoad()

	// Apply saved settings
	var settings api.SettingsResponse
	settingsErr := mapstructure.Decode(hp.Config.Get("settings"), &settings)
	if settingsErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": settingsErr.Error(),
		}).Panic("main - settings load failed")
	}

	applyErr := api.ApplySettings(settings, true)
	if applyErr != nil {
		if e, ok := errors.AsType[*api.Error](applyErr); !ok || e.Code == 500 {
			logrus.WithFields(logrus.Fields{
				"err": applyErr.Error(),
			}).Panic("main - settings apply failed")
		}
	}

	logrus.Info("main - nmcli loaded")
}
