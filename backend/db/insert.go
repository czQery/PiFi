package db

import (
	"strings"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite/sqlitex"
)

func InsertAP(bssid, ssid, mode, discovered string, channel, frequency int64, rssi, latitude, longitude float64, altitude, accuracy int64, device string) {
	err := sqlitex.Execute(Conn, "INSERT INTO ap (bssid, ssid, mode, discovered, channel, frequency,  rssi, latitude, longitude, altitude, accuracy, device) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", &sqlitex.ExecOptions{
		Args: []interface{}{
			bssid,
			ssid,
			mode,
			discovered,
			channel,
			frequency,
			rssi,
			latitude,
			longitude,
			altitude,
			accuracy,
			device,
		},
	})
	if err != nil && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - insert ap exec failed")
	}
}
