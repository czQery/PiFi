package hp

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
	"os"
)

var Dist bool

func LoadDist() {
	_, err := os.ReadDir("./dist")
	if err != nil {
		Dist = false
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("dist - load failed")
		return
	}

	Dist = true
	log.Info("dist - successfully loaded")
}
