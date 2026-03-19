package cmd

import (
	"context"
	"os/exec"
	"strings"

	"github.com/tidwall/gjson"
)

type Interface struct {
	Name        string
	Type        string
	State       string
	Description string
}

func SetRegion(ctx context.Context, region string) error {
	return exec.CommandContext(ctx, "iw", "reg", "set", region).Run()
}

func GetInterfaceList(ctx context.Context) ([]Interface, error) {
	var list []Interface
	out, err := exec.CommandContext(ctx, NM, "-t", "device").Output()
	if err != nil {
		return list, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		i := strings.Split(line, ":")
		if len(i) < 4 {
			continue
		}

		list = append(list, Interface{Name: i[0], Type: i[1], State: i[2], Description: i[3]})
	}

	return list, nil
}

func GetInterfaceBSSID(ctx context.Context, iface string) (string, error) {
	out, err := exec.CommandContext(ctx, "ip", "-j", "link", "show", iface).Output()
	if err != nil {
		return "", err
	}

	return strings.ToLower(gjson.Parse(string(out)).Get(`0.address`).String()), nil
}

func SetInterfaceMode(ctx context.Context, iface string, mode string) error {
	out, err := exec.CommandContext(ctx, "bash", "-c", "iw dev "+iface+" info | grep type").Output()
	modeOld := strings.TrimSpace(string(out))
	modeOld = strings.ReplaceAll(modeOld, " ", "")
	modeOld = strings.ReplaceAll(modeOld, "type", "")
	if modeOld == mode || mode == "managed" && modeOld == "AP" {
		return nil
	}

	if err = exec.CommandContext(ctx, "ip", "link", "set", iface, "down").Run(); err != nil {
		return err
	}

	if err = exec.CommandContext(ctx, "iw", "dev", iface, "set", "type", mode).Run(); err != nil {
		return err
	}

	if err = exec.CommandContext(ctx, "ip", "link", "set", iface, "up").Run(); err != nil {
		return err
	}

	return nil
}
