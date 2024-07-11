package api

import (
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
	"os"
)

func Log(c *fiber.Ctx) error {
	data, err := os.ReadFile(hp.LogFileName)
	if err != nil {
		return &Error{Code: 500, Func: "api/log", Err: err}
	}

	return c.Status(200).JSON(Response{Message: "success", Data: data})
}
