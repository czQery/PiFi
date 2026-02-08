package api

import (
	"strconv"
	"strings"

	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

func WigleGet(c *fiber.Ctx) error {
	c.Context().SetContentType("text/csv")
	payload, _ := net.WigleGet(c.Context(), false)
	return c.SendString(payload)
}

func WiglePatch(c *fiber.Ctx) error {
	value, _ := strconv.ParseBool(c.Query("value"))
	list := strings.Split(c.Query("list"), ",")

	// just mark all valid aps if the user didnt specify the BSSID list
	if c.Query("list") == "" {
		_, list = net.WigleGet(c.Context(), true)
	}

	net.WigleMark(c.Context(), value, list)

	return c.Status(200).JSON(Response{Message: "success"})
}

func WiglePost(c *fiber.Ctx) error {
	err := net.WigleUpload(c.Context())
	if err != nil {
		return &Error{Code: 500, Func: "api/net", Err: err, Message: "upload failed"}
	}

	return c.Status(200).JSON(Response{Message: "success"})
}
