package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Portal(c *fiber.Ctx) error {
	logrus.WithFields(logrus.Fields{
		"data": c.Queries(),
	}).Info("portal - received")

	return c.Status(200).JSON(Response{Message: "success"})
}
