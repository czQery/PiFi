package api

import (
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Log(c *fiber.Ctx) error {
	if !VerifyToken(c) {
		return c.Status(401).JSON(Response{Message: "unauthorized"})
	}

	data, err := os.ReadFile(hp.LogFileName)
	if err != nil {
		return c.Status(500).JSON(Response{Message: "error"})
	}

	return c.Status(200).JSON(Response{Message: "success", Data: data})
}
