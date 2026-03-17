package hp

import (
	"os"
	"strings"
	"time"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/sirupsen/logrus"
)

var Config = koanf.New(".")

const ConfigName = "config.toml"

func ConfigLoad() {
	err := Config.Load(file.Provider(ConfigName), toml.Parser())
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			err = os.WriteFile(ConfigName, []byte("[main]\naddress = \":80\"\npassword = \"ligma\"\ngateway = \"10.42.0.1\"\nregion = \"CZ\""), 0644)
			if err == nil {
				logrus.Info("config - created default config")
				time.Sleep(3 * time.Second) // prevent very fast infinity loop just in case something went wrong
				ConfigLoad()
				return
			}
		}

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
		return
	}

	err = os.WriteFile(ConfigName, data, 0644)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("config - save failed")
	}
}

func ConfigGetInterfaceList() map[string]map[string]interface{} {
	ifaceConfig := make(map[string]map[string]interface{})
	ifaceConfigErr := Config.Unmarshal("settings.iface", &ifaceConfig)
	if ifaceConfigErr != nil {
		logrus.WithFields(logrus.Fields{
			"err": ifaceConfigErr.Error(),
		}).Panic("hp - iface config load failed")
	}

	return ifaceConfig
}
