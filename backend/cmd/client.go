package cmd

import (
	"context"
	"errors"
	"os/exec"
)

func SetClient(ctx context.Context, iface, ssid, password string) error {
	err := exec.CommandContext(ctx, "device", "wifi", "rescan", "ifname", iface).Run()
	if err != nil {
		return errors.New("wifi rescan: " + err.Error())
	}

	_ = exec.CommandContext(ctx, NM, "con", "delete", Con+"-client").Run()

	err = exec.CommandContext(ctx, NM, "device", "wifi", "connect", ssid, "password", password, "ifname", iface, "hidden", "yes", "name", Con+"-client").Run()
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

func DisableClient(ctx context.Context) error {
	return exec.CommandContext(ctx, NM, "con", "delete", Con+"-client").Run()
}
