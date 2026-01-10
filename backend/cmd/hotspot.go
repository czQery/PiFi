package cmd

import (
	"errors"
	"os/exec"
	"strings"
)

func InitHotspot(iface string, item map[string]interface{}) (map[string]interface{}, error) {
	out, err := exec.Command(NM, "-t", "con").Output()
	if err != nil {
		return item, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		c := strings.Split(line, ":")
		if len(c) < 4 {
			continue
		}

		if c[0] == Con+"-hotspot" {
			item["mode"] = "hotspot"
			return item, nil // hotspot already initialized
		}
	}

	err = exec.Command(NM, "con", "add", "type", "wifi", "ifname", iface, "con-name", Con+"-hotspot", "autoconnect", "yes", "ssid", Con).Run()
	if err != nil {
		return item, err
	}

	err = exec.Command(NM, "con", "modify", Con+"-hotspot", "802-11-wireless.mode", "ap", "802-11-wireless.band", "bg", "802-11-wireless.channel", "1", "ipv4.method", "shared").Run()
	if err != nil {
		return item, err
	}

	err = exec.Command(NM, "con", "up", Con+"-hotspot").Run()
	if err != nil {
		return item, err
	}

	item["mode"] = "hotspot"
	item["ssid"] = Con
	item["channel"] = 1

	return item, nil
}

func SetHotspot(iface, ssid, channel, password string) error {
	err := exec.Command(NM, "con", "modify", Con+"-hotspot", "connection.interface-name", iface, "802-11-wireless.ssid", ssid, "802-11-wireless.band", "bg", "802-11-wireless.channel", channel).Run()
	if err != nil {
		return errors.New("modify iface: " + err.Error())
	}

	if password == "" || len(password) < 8 {
		err = exec.Command(NM, "con", "modify", Con+"-hotspot", "remove", "802-11-wireless-security").Run()
	} else {
		err = exec.Command(NM, "con", "modify", Con+"-hotspot", "802-11-wireless-security.key-mgmt", "wpa-psk", "802-11-wireless-security.psk", password, "802-11-wireless-security.pmf", "disable").Run()
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

	err = exec.Command(NM, "con", "up", Con+"-hotspot").Run()
	if err != nil {
		return errors.New("up: " + err.Error())
	}

	return nil
}

func DisableHotspot() error {
	return exec.Command(NM, "con", "down", Con+"-hotspot").Run()
}
