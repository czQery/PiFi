package api

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {

	if !VerifyToken(c) {
		return &Error{Code: 401, Func: "api/auth", Message: "unauthorized"}
	}

	return c.Status(200).JSON(Response{Message: "success"})
}

func VerifyToken(c *fiber.Ctx) bool {
	hash := sha256.Sum256([]byte(hp.Config.String("main.password")))

	if string(c.Request().Header.Cookie("token")) == hex.EncodeToString(hash[:]) {
		return true
	}

	return false
}
