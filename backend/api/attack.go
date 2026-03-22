package api

import (
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var DeauthTargets sync.Map
var ChannelSwitchTargets sync.Map

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

func ChannelSwitchGet(c *fiber.Ctx) error {
	var list = make(map[string]int)
	ChannelSwitchTargets.Range(func(key, value interface{}) bool {
		list[key.(string)] = value.(int)
		return true
	})
	return c.Status(200).JSON(Response{Message: "success", Data: list})
}

func ChannelSwitchPost(c *fiber.Ctx) error {
	target := strings.ToLower(c.Query("target"))
	channel, _ := strconv.Atoi(c.Query("channel"))

	if channel < 1 || channel > 13 {
		return c.Status(400).JSON(Response{Message: "channel out of range"})
	}

	if target == "ff:ff:ff:ff:ff:ff" || target == "" || target == "*" {
		return c.Status(400).JSON(Response{Message: "ff:ff:ff:ff:ff:ff not allowed"})
	}

	ChannelSwitchTargets.Store(c.Query("target"), channel)
	return c.Status(200).JSON(Response{Message: "success"})
}

func ChannelSwitchDelete(c *fiber.Ctx) error {
	ChannelSwitchTargets.Delete(c.Query("target"))
	return c.Status(200).JSON(Response{Message: "success"})
}
