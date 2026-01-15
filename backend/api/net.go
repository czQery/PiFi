package api

import (
	"strconv"
	"strings"

	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

func WigleGet(c *fiber.Ctx) error {
	c.Context().SetContentType("text/csv")
	payload, _ := net.WigleGet()
	return c.SendString(payload)
}

func WiglePatch(c *fiber.Ctx) error {

	value, _ := strconv.ParseBool(c.Query("value"))
	list := strings.Split(c.Query("list"), ",")

	net.WigleMark(value, list)

	return c.Status(200).JSON(Response{Message: "success"})
}

func WiglePost(c *fiber.Ctx) error {

	err := net.WigleUpload()
	if err != nil {
		return &Error{Code: 500, Func: "api/net", Err: err, Message: "upload failed"}
	}

	return c.Status(200).JSON(Response{Message: "success"})
}
