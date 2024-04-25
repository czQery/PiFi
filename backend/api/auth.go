package api

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	token := string(c.Request().Header.Cookie("token"))
	hash := sha256.Sum256([]byte(hp.Config.String("main.password")))

	if token == hex.EncodeToString(hash[:]) {
		return c.Status(200).JSON(Response{Message: "success"})
	}

	return c.Status(401).JSON(Response{Message: "unauthorized"})
}
