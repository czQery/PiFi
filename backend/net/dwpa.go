package net

import (
	"context"
	"os"
	"strings"

	"github.com/czQery/PiFi/backend/db"
)

type Handshake struct {
	File  string `json:"file"`
	BSSID string `json:"bssid"`
	SSID  string `json:"ssid"`
	Net   bool   `json:"net"`
}

func DWPAGet(ctx context.Context) []Handshake {
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

	list := db.SelectAPInList(ctx, "bssid", bssids)
	var handshakes []Handshake
	for _, ap := range list {
		for i, bssid := range bssids {
			if bssid == ap.BSSID {
				handshakes = append(handshakes, Handshake{File: files[i], BSSID: ap.BSSID, SSID: ap.SSID, Net: ap.DWPA})
			}
		}
	}

	return handshakes
}

func DWPAUpload(ctx context.Context) error {
	return nil
}
