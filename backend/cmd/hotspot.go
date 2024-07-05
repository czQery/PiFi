package cmd

import (
	"os/exec"
	"strings"
)

func InitHotspot(ifname string) error {
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

	err = exec.Command(NM, "con", "add", "type", "wifi", "ifname", ifname, "con-name", Con, "autoconnect", "yes", "ssid", "PiFi").Run()
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

func SetHotspot(ifname, ssid, channel, password string) error {

	err := exec.Command(NM, "con", "modify", Con, "connection.interface-name", ifname, "802-11-wireless.ssid", ssid, "802-11-wireless.channel", channel).Run()
	if err != nil {
		return err
	}

	if password != "" {
		err = exec.Command(NM, "con", "modify", Con, "802-11-wireless-security.key-mgmt", "wpa-psk", "802-11-wireless-security.psk", "'"+password+"'").Run()
	} else {
		err = exec.Command(NM, "con", "modify", Con, "remove", "802-11-wireless-security").Run()
	}
	if err != nil {
		return err
	}

	err = exec.Command(NM, "con", "up", Con).Run()
	if err != nil {
		return err
	}

	return nil
}
