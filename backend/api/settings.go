package api

import (
	"encoding/json"
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

type SettingsResponse struct {
	Interface map[string]SettingsInterfaceResponse `json:"iface"`
}

type SettingsInterfaceResponse struct {
	Mode     string `json:"mode"`
	Ready    bool   `json:"ready"`
	SSID     string `json:"ssid"`
	Password string `json:"password"`
	Channel  int    `json:"channel"`
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

	err = ApplySettings(data)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	return SettingsGet(c)
}

func ApplySettings(settings SettingsResponse) error {

	var hotspot bool

	for ifaceName, iface := range settings.Interface {
		if !iface.Ready {
			continue
		}

		switch strings.ToLower(iface.Mode) {
		case "hotspot":
			err := cmd.SetHotspot(ifaceName, iface.SSID, strconv.Itoa(iface.Channel), iface.Password)
			if err != nil {
				return err
			}
			hotspot = true
		}
	}

	if !hotspot {
		err := cmd.DisableHotspot()
		if err != nil {
			return err
		}
	}

	// very retarded approach, but it works I guess
	// struct -> json -> map (map because toml parser likes it)
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
