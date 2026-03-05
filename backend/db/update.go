package db

import (
	"context"
	"encoding/json"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite/sqlitex"
)

func UpdateAPNet(ctx context.Context, net string, value bool, list []string) {
	conn, err := Pool.Take(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap net connection failed")
		return
	}
	defer Pool.Put(conn)

	listJson, _ := json.Marshal(list)
	err = sqlitex.Execute(conn, "UPDATE ap SET "+net+" = ? WHERE bssid IN (SELECT value FROM json_each(?));", &sqlitex.ExecOptions{
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

func UpdateAPHandshake(ctx context.Context, bssid string, value bool) {
	conn, err := Pool.Take(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap handshake connection failed")
		return
	}
	defer Pool.Put(conn)

	err = sqlitex.Execute(conn, "UPDATE ap SET handshake = ? WHERE bssid = ?;", &sqlitex.ExecOptions{
		Args: []interface{}{
			value,
			bssid,
		},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap handshake exec failed")
	}
}

func UpdateAP(ctx context.Context, bssid, ssid, mode string, channel, frequency int64, rssi, latitude, longitude, altitude, accuracy float64) {
	conn, err := Pool.Take(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - update ap connection failed")
		return
	}
	defer Pool.Put(conn)

	err = sqlitex.Execute(conn, "UPDATE ap SET ssid = ?, mode = ?, channel = ?, frequency = ?, rssi = ?, latitude = ?, longitude = ?, altitude = ?, accuracy = ? WHERE bssid = ?;", &sqlitex.ExecOptions{
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
