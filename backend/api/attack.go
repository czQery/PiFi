package api

import (
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var DeauthTargets sync.Map

func DeauthGet(c *fiber.Ctx) error {
	var list []string
	DeauthTargets.Range(func(key, _ interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	return c.Status(200).JSON(Response{Message: "success", Data: list})
}

func DeauthPost(c *fiber.Ctx) error {
	target := strings.ToLower(c.Query("target"))

	if target == "ff:ff:ff:ff:ff:ff" || target == "" || target == "*" {
		return c.Status(400).JSON(Response{Message: "ff:ff:ff:ff:ff:ff not allowed"})
	}

	DeauthTargets.Store(c.Query("target"), struct{}{})
	return c.Status(200).JSON(Response{Message: "success"})
}

func DeauthDelete(c *fiber.Ctx) error {
	DeauthTargets.Delete(c.Query("target"))
	return c.Status(200).JSON(Response{Message: "success"})
}
