package api

import (
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
	DeauthTargets.Store(c.Query("target", "ff:ff:ff:ff:ff:ff"), struct{}{})
	return c.Status(200).JSON(Response{Message: "success"})
}

func DeauthDelete(c *fiber.Ctx) error {
	DeauthTargets.Delete(c.Query("target", "ff:ff:ff:ff:ff:ff"))
	return c.Status(200).JSON(Response{Message: "success"})
}
