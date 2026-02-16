package net

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func DWPAGet() []string {
	dir, dirErr := os.ReadDir("./cap")
	if dirErr != nil {
		return nil
	}

	var files []string
	var bssids []string

	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil || info.Size() <= 128 { // skip empty files
			continue
		}

		split := strings.Split(info.Name(), "_")
		if len(split) <= 1 {
			continue
		}

		bssid := strings.TrimSuffix(split[1], ".pcap")
		if len(bssid)%2 != 0 {
			continue
		}

		var octets []string
		for i := 0; i < len(bssid); i += 2 {
			octets = append(octets, bssid[i:i+2])
		}

		files = append(files, entry.Name())
		bssids = append(bssids, strings.Join(octets, ":"))
	}

	fmt.Println(bssids)

	return files
}

func DWPAUpload(ctx context.Context) error {
	return nil
}
