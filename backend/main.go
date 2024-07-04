package main

import (
	"encoding/json"
	"errors"
	"github.com/czQery/PiFi/backend/api"
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
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

	hp.ArtPrint()
	hp.ConfigLoad()
	hp.DistLoad()

	// Get interfaces
	iface, ifaceErr := cmd.GetInterfaceList()
	if ifaceErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceErr.Error(),
		}).Panic("main - " + cmd.NM + " failed")
	}

	// Create hotspot con for the first time
	initErr := errors.New("no wifi interface")
	for _, i := range iface {
		if i.Type == "wifi" {
			initErr = cmd.InitHotspot(i.Name)
			break
		}
	}
	if initErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": initErr.Error(),
		}).Panic("main - hotspot init failed")
	}

	logrus.Info("main - " + cmd.NM + " loaded")

	r := fiber.New(fiber.Config{
		CaseSensitive:         false,
		DisableStartupMessage: true,
		GETOnly:               false,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		ServerHeader:          "PiFi",
	})

	// API
	r.Get("/api/auth", api.Auth)
	r.Get("/api/stats", api.Stats)
	r.Get("/api/log", api.Log)

	// Static files
	r.Static("/", "./dist")

	// Default
	r.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(string(c.Request().URI().Path()), "/api") {
			return c.Status(404).JSON(api.Response{Message: "unknown endpoint"})
		} else if !hp.Dist {
			return c.Status(503).JSON(api.Response{Message: "front-end unavailable"})
		} else {
			return c.Redirect("/", 307)
		}
	})

	logrus.Info("fiber - started")

	// Run
	err := r.Listen(hp.Config.String("main.address"))
	logrus.WithFields(logrus.Fields{
		"err": err.Error(),
	}).Panic("fiber - server failed")

	_ = hp.LogFile.Close()
}
