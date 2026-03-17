package api

import (
	"context"
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
	Main      SettingsMainResponse                 `json:"main"`
	Hotspot   SettingsHotspotResponse              `json:"hotspot"`
	Client    SettingsClientResponse               `json:"client"`
	Monitor   SettingsMonitorResponse              `json:"monitor"`
	Interface map[string]SettingsInterfaceResponse `json:"iface" mapstructure:"iface" koanf:"iface"`
}

type SettingsMainResponse struct {
	Address string `json:"address"`
	Gateway string `json:"gateway"`
	Region  string `json:"region"`
}

type SettingsHotspotResponse struct {
	SSID         string `json:"ssid"`
	Password     string `json:"password"`
	Channel      int    `json:"channel"`
	Portal       bool   `json:"portal"`
	PortalSource string `json:"portal_source" mapstructure:"portal_source" koanf:"portal_source"`
}

type SettingsClientResponse struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
}

type SettingsMonitorResponse struct {
	Auto bool `json:"auto"`
}

type SettingsInterfaceResponse struct {
	Mode  string `json:"mode"`
	Ready bool   `json:"ready"`
}

func SettingsGet(c *fiber.Ctx) error {
	var settings SettingsResponse
	err := hp.Config.Unmarshal("settings", &settings)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: settings})
}

func SettingsPost(c *fiber.Ctx) error {
	var data SettingsResponse
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err, Message: "invalid body"}
	}

	err = ApplySettings(c.Context(), data, false)
	if err != nil {
		return err
	}

	return SettingsGet(c)
}

func ApplySettings(ctx context.Context, settings SettingsResponse, force bool) error {

	var (
		err error

		modes     = make(map[string]struct{})
		modesLast = make(map[string]struct{})
	)

	var settingsSaved SettingsResponse
	configErr := mapstructure.Decode(hp.Config.Get("settings"), &settingsSaved)
	if configErr != nil {
		return errors.New("config read: " + configErr.Error())
	}

	for _, iface := range settingsSaved.Interface {
		modesLast[iface.Mode] = struct{}{}
	}

	for _, iface := range settings.Interface {
		if _, ok := modes[iface.Mode]; ok {
			return &Error{Code: 400, Func: "api/settings", Message: "duplicit mode"}
		}

		if iface.Mode == "none" {
			continue
		}

		modes[iface.Mode] = struct{}{}
	}

	for ifaceName, iface := range settings.Interface {
		if !iface.Ready {
			continue
		}

		canSkip := !force && settingsSaved.Interface[ifaceName] == iface

		switch strings.ToLower(iface.Mode) {
		case "hotspot":
			if canSkip && settings.Hotspot == settingsSaved.Hotspot {
				continue
			}

			err = cmd.SetInterfaceMode(ctx, ifaceName, "managed")
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"iface": ifaceName,
					"mode":  iface.Mode,
					"err":   err.Error(),
				}).Error("cmd - failed to set interface mode")
			}

			if settings.Hotspot.Portal {
				cmd.HotspotPortal = settings.Hotspot.PortalSource
			} else {
				cmd.HotspotPortal = ""
			}

			if settings.Hotspot.SSID == "" {
				settings.Hotspot.SSID = cmd.Con
			}
			if settings.Hotspot.Channel == 0 {
				settings.Hotspot.Channel = 1
			}

			cmd.Hotspot = settings.Hotspot.SSID
			cmd.HotspotBSSID, err = cmd.GetInterfaceBSSID(ctx, ifaceName)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"iface": ifaceName,
					"err":   err.Error(),
				}).Warn("cmd - failed to get BSSID of hotspot")
			}

			logrus.WithFields(logrus.Fields{
				"iface":    ifaceName,
				"ssid":     settings.Hotspot.SSID,
				"channel":  settings.Hotspot.Channel,
				"password": settings.Hotspot.Password,
				"portal":   cmd.HotspotPortal,
			}).Info("cmd - setting up hotspot")
			err = cmd.SetHotspot(ctx, ifaceName, settings.Hotspot.SSID, strconv.Itoa(settings.Hotspot.Channel), settings.Hotspot.Password, settings.Hotspot.Portal)
			if err != nil && !force {
				return errors.New("set hotspot: " + err.Error())
			}
		case "client":
			if canSkip && settings.Client == settingsSaved.Client {
				continue
			}

			err = cmd.SetInterfaceMode(ctx, ifaceName, "managed")
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"iface": ifaceName,
					"mode":  iface.Mode,
					"err":   err.Error(),
				}).Error("cmd - failed to set interface mode")
			}

			logrus.WithFields(logrus.Fields{
				"iface":    ifaceName,
				"ssid":     settings.Client.SSID,
				"password": settings.Client.Password,
			}).Info("cmd - connecting to wifi")
			err = cmd.SetClient(ctx, ifaceName, settings.Client.SSID, settings.Client.Password)
			if err != nil && !force {
				logrus.WithFields(logrus.Fields{
					"iface":    ifaceName,
					"ssid":     settings.Client.SSID,
					"password": settings.Client.Password,
				}).Info("cmd - connecting to fallback wifi")
				_ = cmd.SetClient(ctx, ifaceName, settings.Client.SSID, settings.Client.Password) // fallback to previously saved wifi
				return &Error{Code: 400, Func: "api/settings/client", Err: err, Message: err.Error()}
			}
		case "monitor":
			if canSkip && settings.Monitor == settingsSaved.Monitor {
				continue
			}

			err = cmd.SetInterfaceMode(ctx, ifaceName, "monitor")
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"iface": ifaceName,
					"mode":  iface.Mode,
					"err":   err.Error(),
				}).Error("cmd - failed to set interface mode")
			}

			cmd.MonitorAuto = settings.Monitor.Auto

			logrus.WithFields(logrus.Fields{
				"iface": ifaceName,
			}).Info("cmd - setting up bettercap monitor")
			_ = cmd.SetBettercap(ctx, "set wifi.region "+settingsSaved.Main.Region)
			err = cmd.SetBettercapMonitor(ctx, ifaceName)
			if err != nil && !force {
				return errors.New("set monitor: " + err.Error())
			}
		}
	}

	if _, ok := modes["hotspot"]; !ok {
		if _, ok = modesLast["hotspot"]; ok {
			logrus.Info("cmd - disabling hotspot")
		}
		cmd.Hotspot = ""
		cmd.HotspotPortal = ""
		_ = cmd.DisableHotspot(ctx)
	}
	if _, ok := modes["client"]; !ok {
		if _, ok = modesLast["client"]; ok {
			logrus.Info("cmd - disabling client")
		}
		_ = cmd.DisableClient(ctx)
	}
	if _, ok := modes["monitor"]; !ok {
		if _, ok = modesLast["monitor"]; ok {
			logrus.Info("cmd - disabling bettercap monitor")
		}
		_ = cmd.DisableBettercapMonitor(ctx)
	}

	// very retarded approach, but it works I guess
	// struct -> JSON -> map (map because toml parser likes it)
	var settingsMap map[string]interface{}
	var settingsJson []byte
	settingsJson, err = json.Marshal(settings)
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
