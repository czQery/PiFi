package api

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"sort"
	"strings"
)

func Portals(c *fiber.Ctx) error {

	var list []string

	// get dir list
	dirs, err := os.ReadDir("./portal")
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		list = append(list, dir.Name())
	}

	// alphabetically sort list
	sort.Slice(list, func(i, j int) bool {
		return strings.ToLower(list[i]) < strings.ToLower(list[j])
	})

	return c.Status(200).JSON(Response{Message: "success", Data: list})
}
