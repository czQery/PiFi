package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func Portal(c *fiber.Ctx) error {

	var data []string

	for name, value := range c.Queries() {
		data = append(data, name+"='"+value+"'")
	}

	logrus.WithFields(logrus.Fields{
		"data": "[" + strings.Join(data, ",") + "]",
	}).Info("portal - received")

	return c.Status(200).JSON(Response{Message: "success"})
}
