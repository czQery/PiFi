package api

import (
	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

func WigleGet(c *fiber.Ctx) error {
	c.Context().SetContentType("text/csv")
	return c.SendString(net.WigleGet())
}

func WiglePost(c *fiber.Ctx) error {

	err := net.WigleUpload()
	if err != nil {
		return &Error{Code: 500, Func: "api/net", Err: err, Message: "upload failed"}
	}

	return c.Status(200).JSON(Response{Message: "success"})
}
