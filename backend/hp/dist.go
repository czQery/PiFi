package hp

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Dist bool

func DistLoad() {
	_, err := os.ReadDir("./dist")
	if err != nil {
		Dist = false
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("dist - load failed")
		return
	}

	Dist = true
	logrus.Info("dist - successfully loaded")
}
