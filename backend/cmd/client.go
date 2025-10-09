package cmd

import (
	"errors"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func SetClient(iface, ssid, password string) error {
	logrus.WithFields(logrus.Fields{
		"iface":    iface,
		"ssid":     ssid,
		"password": password,
	}).Info("cmd - connecting to wifi")

	err := exec.Command(NM, "device", "wifi", "rescan", "ifname", iface).Run()
	if err != nil {
		return errors.New("wifi rescan: " + err.Error())
	}

	_ = exec.Command(NM, "con", "delete", Con+"-client").Run()

	err = exec.Command(NM, "device", "wifi", "connect", ssid, "password", password, "ifname", iface, "hidden", "yes", "name", Con+"-client").Run()
	if err != nil {
		switch err.Error() {
		case "exit status 4":
			return errors.New("wifi connect: wrong password")
		case "exit status 10":
			return errors.New("wifi connect: no such wifi")
		default:
			return errors.New("wifi connect: " + err.Error())
		}
	}

	return nil
}
