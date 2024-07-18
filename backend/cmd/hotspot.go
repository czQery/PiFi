package cmd

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os/exec"
	"strings"
)

func InitHotspot(iface string) error {
	out, err := exec.Command(NM, "-t", "con").Output()
	if err != nil {
		return err
	}

	for _, line := range strings.Split(string(out), "\n") {
		c := strings.Split(line, ":")
		if len(c) < 4 {
			continue
		}

		if c[0] == "PiFi" {
			return nil // hotspot already initialized
		}
	}

	err = exec.Command(NM, "con", "add", "type", "wifi", "ifname", iface, "con-name", Con, "autoconnect", "yes", "ssid", "PiFi").Run()
	if err != nil {
		return err
	}

	err = exec.Command(NM, "con", "modify", Con, "802-11-wireless.mode", "ap", "802-11-wireless.band", "bg", "802-11-wireless.channel", "1", "ipv4.method", "shared").Run()
	if err != nil {
		return err
	}

	err = exec.Command(NM, "con", "up", Con).Run()
	if err != nil {
		return err
	}

	return nil
}

func SetHotspot(iface, ssid, channel, password string) error {
	logrus.WithFields(logrus.Fields{
		"iface":    iface,
		"ssid":     ssid,
		"channel":  channel,
		"password": password,
		"portal":   Portal != "",
	}).Info("cmd - setting up hotspot")

	err := exec.Command(NM, "con", "modify", Con, "connection.interface-name", iface, "802-11-wireless.ssid", ssid, "802-11-wireless.channel", channel).Run()
	if err != nil {
		return errors.New("modify iface: " + err.Error())
	}

	if password == "" || len(password) < 8 {
		err = exec.Command(NM, "con", "modify", Con, "remove", "802-11-wireless-security").Run()
	} else {
		err = exec.Command(NM, "con", "modify", Con, "802-11-wireless-security.key-mgmt", "wpa-psk", "802-11-wireless-security.psk", password).Run()
	}
	if err != nil {
		return errors.New("modify wpa: " + err.Error())
	}

	if Portal == "" {
		err = DisableDNSPortal()
	} else {
		err = SetDNSPortal()
	}
	if err != nil {
		return errors.New("portal dns: " + err.Error())
	}

	err = exec.Command(NM, "con", "up", Con).Run()
	if err != nil {
		return errors.New("up: " + err.Error())
	}

	return nil
}

func DisableHotspot() error {
	logrus.Info("cmd - disabling hotspot")
	return exec.Command(NM, "con", "down", Con).Run()
}
