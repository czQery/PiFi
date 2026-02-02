package cmd

import (
	"errors"
	"os/exec"
	"strings"
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

func SetInterfaceMode(iface string, mode string) error {
	out, err := exec.Command("ip", "link", "set", iface, "down").Output()
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
