package hp

import (
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/sirupsen/logrus"
)

var Config = koanf.New(".")

func ConfigLoad() {
	err := Config.Load(file.Provider("config.toml"), toml.Parser())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Panic("config - load failed")
	}

	logrus.Info("config - successfully loaded")
}
