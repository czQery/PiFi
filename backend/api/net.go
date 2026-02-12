package api

import (
	"strconv"
	"strings"

	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/net"
	"github.com/gofiber/fiber/v2"
)

func NetPatch(c *fiber.Ctx) error {
	netParam := c.Params("net")
	value, _ := strconv.ParseBool(c.Query("value"))
	list := strings.Split(c.Query("list"), ",")

	// just mark all valid aps if the user didnt specify the BSSID list
	if c.Query("list") == "" {
		switch netParam {
		case "wigle":
			_, list = net.WigleGet(c.Context(), true)
		case "beacondb":
		case "dwpa":
		}
	}

	db.UpdateAPNet(c.Context(), c.Query("net"), value, list)
	return c.Status(200).JSON(Response{Message: "success"})
}

func NetPost(c *fiber.Ctx) error {
	var err error

	switch c.Params("net") {
	case "wigle":
		err = net.WigleUpload(c.Context())
	case "beacondb":
	case "dwpa":
		err = net.DWPAUpload(c.Context())
	}

	if err != nil {
		return &Error{Code: 500, Func: "api/net", Err: err, Message: "upload failed"}
	}

	return c.Status(200).JSON(Response{Message: "success"})
}

func WigleGet(c *fiber.Ctx) error {
	c.Context().SetContentType("text/csv")
	payload, _ := net.WigleGet(c.Context(), false)
	return c.SendString(payload)
}

func DWPAGet(c *fiber.Ctx) error {
	return c.Status(200).JSON(Response{Message: "success", Data: net.DWPAGet()})
}
