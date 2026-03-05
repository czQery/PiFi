package db

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite/sqlitex"
)

func InsertAP(ctx context.Context, bssid, ssid, mode, discovered string, channel, frequency int64, rssi, latitude, longitude, altitude, accuracy float64, device string) {
	conn, err := Pool.Take(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - insert ap connection failed")
		return
	}
	defer Pool.Put(conn)

	err = sqlitex.Execute(conn, "INSERT INTO ap (bssid, ssid, mode, discovered, channel, frequency,  rssi, latitude, longitude, altitude, accuracy, device) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", &sqlitex.ExecOptions{
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
				UpdateAP(context.Background(), bssid, ssid, mode, channel, frequency, rssi, latitude, longitude, altitude, accuracy)
			}
			return
		}

		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - insert ap exec failed")
		return
	}

	logrus.WithFields(logrus.Fields{
		"bssid": bssid,
		"ssid":  ssid,
		"rssi":  rssi,
	}).Info("db - new ap")
}
