package db

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite/sqlitex"
)

func UpdateAPNet(net string, value bool, list []string) {
	listJson, _ := json.Marshal(list)
	err := sqlitex.Execute(Conn, "UPDATE ap SET "+net+" = ? WHERE bssid IN (SELECT value FROM json_each(?));", &sqlitex.ExecOptions{
		Args: []interface{}{
			value,
			listJson,
		},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap net exec failed")
	}
}

func UpdateAP(bssid, ssid, mode string, channel, frequency int64, rssi, latitude, longitude, altitude, accuracy float64) {
	err := sqlitex.Execute(Conn, "UPDATE ap SET ssid = ?, mode = ?, channel = ?, frequency = ?, rssi = ?, latitude = ?, longitude = ?, altitude = ?, accuracy = ? WHERE bssid = ?;", &sqlitex.ExecOptions{
		Args: []interface{}{
			ssid,
			mode,
			channel,
			frequency,
			rssi,
			latitude,
			longitude,
			altitude,
			accuracy,
			bssid,
		},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap exec failed")
	}
}
