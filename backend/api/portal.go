package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Portal(c *fiber.Ctx) error {

	var data []string

	for name, value := range c.Queries() {
		data = append(data, name+"='"+value+"'")
	}

	if len(data) != 0 {
		logrus.WithFields(logrus.Fields{
			"data": "[" + strings.Join(data, ",") + "]",
		}).Info("portal - received")
	}

	return c.Status(200).JSON(Response{Message: "success"})
}
