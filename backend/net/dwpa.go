package net

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
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
		if err != nil || info.Size() <= 64 { // skip empty files
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
			if bssid == ap.BSSID && ap.Handshake {
				handshakes = append(handshakes, Handshake{File: files[i], BSSID: ap.BSSID, SSID: ap.SSID, Net: ap.DWPA})
			}
		}
	}

	return handshakes
}

func DWPAUpload(ctx context.Context) error {
	if hp.Config.Get("net.dwpa") == nil {
		return errors.New("upload failed: missing dwpa key")
	}

	key := hp.Config.Get("net.dwpa").(string)

	logrus.WithFields(logrus.Fields{
		"key": key,
	}).Debug("net - dwpa upload")

	list := DWPAGet(ctx)
	cl := req.NewClient()
	var bssids []string
	for _, ap := range list {
		if ap.Net {
			continue
		}

		up, upErr := cl.R().SetHeader("Cookie", "key="+key).SetFile("file", "./cap/"+ap.File).Post("https://wpa-sec.stanev.org/?submit=")
		if upErr != nil {
			return upErr
		}

		if up.StatusCode != 200 {
			return errors.New("upload failed: code " + strconv.Itoa(up.StatusCode))
		}

		bssids = append(bssids, ap.BSSID)

		logrus.WithFields(logrus.Fields{
			"file": ap.File,
		}).Info("net - dwpa successfully uploaded")
	}

	db.UpdateAPNet(ctx, "dwpa", true, bssids)
	return nil
}
