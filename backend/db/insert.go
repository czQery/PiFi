package db

import (
	"strings"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite/sqlitex"
)

func InsertAP(bssid, ssid, mode, discovered string, channel, frequency int64, rssi, latitude, longitude, altitude, accuracy float64, device string) {
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
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			if accuracy != 0 || longitude != 0 || latitude != 0 { // automatically update record if the new data has valid location
				UpdateAP(bssid, ssid, mode, channel, frequency, rssi, latitude, longitude, altitude, accuracy)
			}
			return
		}

		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - insert ap exec failed")
	}
}
