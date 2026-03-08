package cmd

import (
	"context"
	"os/exec"
	"strings"
)

type Connection struct {
	Name      string
	UUID      string
	Type      string
	Interface string
}

func GetConnectionList(ctx context.Context) ([]Connection, error) {
	var list []Connection

	out, err := exec.CommandContext(ctx, NM, "-t", "con").Output()
	if err != nil {
		return list, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		c := strings.Split(line, ":")
		if len(c) < 4 {
			continue
		}

		list = append(list, Connection{
			Name:      c[0],
			UUID:      c[1],
			Type:      c[2],
			Interface: c[3],
		})
	}

	return list, nil
}
