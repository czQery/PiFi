package cmd

import (
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
