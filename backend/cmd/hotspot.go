package cmd

import (
	"context"
	"errors"
	"os/exec"
)

func InitHotspot(ctx context.Context, iface string, item map[string]interface{}) (map[string]interface{}, map[string]interface{}, error) {
	hotspot := map[string]interface{}{
		"ssid":    Con,
		"channel": 1,
	}

	err := exec.CommandContext(ctx, NM, "con", "add", "type", "wifi", "ifname", iface, "con-name", Con+"-hotspot", "autoconnect", "yes", "ssid", Con).Run()
	if err != nil {
		return item, hotspot, err
	}

	err = exec.CommandContext(ctx, NM, "con", "modify", Con+"-hotspot", "802-11-wireless.mode", "ap", "802-11-wireless.band", "bg", "802-11-wireless.channel", "1", "ipv4.method", "shared").Run()
	if err != nil {
		return item, hotspot, err
	}

	err = exec.CommandContext(ctx, NM, "con", "up", Con+"-hotspot").Run()
	if err != nil {
		return item, hotspot, err
	}

	item["mode"] = "hotspot"

	return item, hotspot, nil
}

func SetHotspot(ctx context.Context, iface, ssid, channel, password string, portal bool) error {
	err := exec.CommandContext(ctx, NM, "con", "modify", Con+"-hotspot", "connection.interface-name", iface, "802-11-wireless.ssid", ssid, "802-11-wireless.band", "bg", "802-11-wireless.channel", channel).Run()
	if err != nil {
		return errors.New("modify iface: " + err.Error())
	}

	if password == "" || len(password) < 8 {
		err = exec.CommandContext(ctx, NM, "con", "modify", Con+"-hotspot", "remove", "802-11-wireless-security").Run()
	} else {
		err = exec.CommandContext(ctx, NM, "con", "modify", Con+"-hotspot", "802-11-wireless-security.key-mgmt", "wpa-psk", "802-11-wireless-security.psk", password, "802-11-wireless-security.pmf", "disable").Run()
	}
	if err != nil {
		return errors.New("modify wpa: " + err.Error())
	}

	if portal {
		err = SetDNSPortal()
	} else {
		err = DisableDNSPortal()
	}
	if err != nil {
		return errors.New("portal dns: " + err.Error())
	}

	err = exec.CommandContext(ctx, NM, "con", "up", Con+"-hotspot").Run()
	if err != nil {
		return errors.New("up: " + err.Error())
	}

	return nil
}

func DisableHotspot(ctx context.Context) error {
	return exec.CommandContext(ctx, NM, "con", "down", Con+"-hotspot").Run()
}
