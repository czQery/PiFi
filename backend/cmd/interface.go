package cmd

import (
	"errors"
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

func GetInterfaceList() ([]Interface, error) {
	var list []Interface
	out, err := exec.Command(NM, "-t", "device").Output()
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

func GetInterfaceBSSID(iface string) (string, error) {
	out, err := exec.Command("ip", "-j", "link", "show", iface).Output()
	if err != nil {
		return "", err
	}

	return strings.ToLower(gjson.Parse(string(out)).Get(`0.address`).String()), nil
}

func SetInterfaceMode(iface string, mode string) error {
	out, err := exec.Command("bash", "-c", "iw dev "+iface+" info | grep type").Output()
	modeOld := strings.TrimSpace(string(out))
	modeOld = strings.ReplaceAll(modeOld, " ", "")
	modeOld = strings.ReplaceAll(modeOld, "type", "")
	if modeOld == mode || mode == "managed" && modeOld == "AP" {
		return nil
	}

	out, err = exec.Command("ip", "link", "set", iface, "down").Output()
	if err != nil {
		return errors.New(string(out))
	}

	out, err = exec.Command("iw", "dev", iface, "set", "type", mode).Output()
	if err != nil {
		return errors.New(string(out))
	}

	out, err = exec.Command("ip", "link", "set", iface, "up").Output()
	if err != nil {
		return errors.New(string(out))
	}

	return nil
}
