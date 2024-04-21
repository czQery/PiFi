package api

import "github.com/gofiber/fiber/v2"

func Auth(c *fiber.Ctx) error {
	return c.Status(401).JSON(Response{Message: "unauthorized"})
}
