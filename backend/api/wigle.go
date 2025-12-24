package api

import (
	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

func WigleGet(c *fiber.Ctx) error {
	c.Context().SetContentType("text/csv")
	return c.SendString(net.WigleGet())
}
