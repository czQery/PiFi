package api

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/go-viper/mapstructure/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type SettingsResponse struct {
	Interface map[string]SettingsInterfaceResponse `json:"iface" mapstructure:"iface" koanf:"iface"`
}

type SettingsInterfaceResponse struct {
	Mode         string `json:"mode"`
	Ready        bool   `json:"ready"`
	SSID         string `json:"ssid"`
	Password     string `json:"password"`
	Channel      int    `json:"channel"`
	Portal       bool   `json:"portal"`
	PortalSource string `json:"portal_source" mapstructure:"portal_source" koanf:"portal_source"`
}

func SettingsGet(c *fiber.Ctx) error {
	iface := make(map[string]SettingsInterfaceResponse)
	err := hp.Config.Unmarshal("settings.iface", &iface)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: SettingsResponse{Interface: iface}})
}

func SettingsPost(c *fiber.Ctx) error {
	var data SettingsResponse
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err, Message: "invalid body"}
	}

	err = ApplySettings(data, false)
	if err != nil {
		return err
	}

	return SettingsGet(c)
}

func ApplySettings(settings SettingsResponse, force bool) error {

	var (
		modeHotspot bool
		modeClient  bool
		modeMonitor bool

		modeHotspotLast bool
		modeClientLast  bool
		modeMonitorLast bool
	)

	for ifaceName, iface := range settings.Interface {
		if !iface.Ready {
			continue
		}

		var config SettingsInterfaceResponse
		configErr := mapstructure.Decode(hp.Config.Get("settings.iface."+ifaceName), &config)
		if configErr != nil {
			return errors.New("config read: " + configErr.Error())
		}

		switch strings.ToLower(config.Mode) {
		case "hotspot":
			modeHotspotLast = true
		case "client":
			modeClientLast = true
		case "monitor":
			modeMonitorLast = true
		}

		switch strings.ToLower(iface.Mode) {
		case "hotspot":
			modeHotspot = true
		case "client":
			modeClient = true
		case "monitor":
			modeMonitor = true
		}

		// skip if the interface is unchanged
		if !force && config == iface {
			continue
		}

		switch strings.ToLower(iface.Mode) {
		case "hotspot":
			if iface.Portal {
				cmd.Portal = iface.PortalSource
			} else {
				cmd.Portal = ""
			}
			logrus.WithFields(logrus.Fields{
				"iface":    ifaceName,
				"ssid":     iface.SSID,
				"channel":  iface.Channel,
				"password": iface.Password,
				"portal":   cmd.Portal,
			}).Info("cmd - setting up hotspot")
			err := cmd.SetHotspot(ifaceName, iface.SSID, strconv.Itoa(iface.Channel), iface.Password)
			if err != nil {
				return errors.New("set hotspot: " + err.Error())
			}
		case "client":
			logrus.WithFields(logrus.Fields{
				"iface":    ifaceName,
				"ssid":     iface.SSID,
				"password": iface.Password,
			}).Info("cmd - connecting to wifi")
			err := cmd.SetClient(ifaceName, iface.SSID, iface.Password)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"iface":    ifaceName,
					"ssid":     config.SSID,
					"password": config.Password,
				}).Info("cmd - connecting to fallback wifi")
				_ = cmd.SetClient(ifaceName, config.SSID, config.Password) // fallback to previously saved wifi
				return &Error{Code: 400, Func: "api/settings/client", Err: err, Message: err.Error()}
			}
		case "monitor":
			logrus.WithFields(logrus.Fields{
				"iface": ifaceName,
			}).Info("cmd - setting up bettercap monitor")
			err := cmd.SetBettercapMonitor(ifaceName)
			if err != nil {
				return errors.New("set monitor: " + err.Error())
			}
		}
	}

	if !modeHotspot {
		if modeHotspotLast {
			logrus.Info("cmd - disabling hotspot")
		}
		_ = cmd.DisableHotspot()
	}
	if !modeClient {
		if modeClientLast {
			logrus.Info("cmd - disabling client")
		}
		_ = cmd.DisableClient()
	}
	if !modeMonitor {
		if modeMonitorLast {
			logrus.Info("cmd - disabling bettercap monitor")
		}
		_ = cmd.DisableBettercapMonitor()
	}

	// very retarded approach, but it works I guess
	// struct -> JSON -> map (map because toml parser likes it)
	var settingsMap map[string]interface{}
	settingsJson, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = json.Unmarshal(settingsJson, &settingsMap)
	if err != nil {
		return err
	}

	err = hp.Config.Set("settings", settingsMap)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	hp.ConfigSave()
	hp.ConfigLoad()

	return nil
}
