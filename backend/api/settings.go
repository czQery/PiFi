package api

import (
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
)

type SettingsResponse struct {
	Interface interface{} `json:"iface"`
}

func SettingsGet(c *fiber.Ctx) error {
	if !VerifyToken(c) {
		return &Error{Code: 401, Func: "api/settings", Message: "unauthorized"}
	}

	iface := make(map[string]map[string]interface{})
	err := hp.Config.Unmarshal("settings.iface", &iface)
	if err != nil {
		return &Error{Code: 500, Func: "api/settings", Err: err}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: SettingsResponse{Interface: iface}})
}
