package hp

import (
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/sirupsen/logrus"
	"os"
)

var Config = koanf.New(".")

const ConfigName = "config.toml"

func ConfigLoad() {
	err := Config.Load(file.Provider(ConfigName), toml.Parser())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("config - load failed")
	}
}

func ConfigSave() {
	data, err := Config.Marshal(toml.Parser())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("config - marshal failed")
	}

	err = os.WriteFile(ConfigName, data, 0644)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("config - save failed")
	}
}
