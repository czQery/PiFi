package db

import (
	"context"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

type DataAP struct {
	BSSID      string  `json:"bssid"`
	SSID       string  `json:"ssid"`
	Mode       string  `json:"mode"`
	Discovered string  `json:"discovered"`
	Channel    int64   `json:"channel"`
	Frequency  int64   `json:"frequency"`
	RSSI       float64 `json:"rssi"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Altitude   float64 `json:"altitude"`
	Accuracy   float64 `json:"accuracy"`
	Device     string  `json:"device"`
	Wigle      bool    `json:"wigle"`
	BeaconDB   bool    `json:"beacondb"`
	DWPA       bool    `json:"dwpa"`
}

func SelectAP(ctx context.Context) []DataAP {
	conn, err := Pool.Take(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - select ap connection failed")
		return nil
	}
	defer Pool.Put(conn)

	var data []DataAP
	err = sqlitex.ExecuteTransient(conn, "SELECT * FROM ap", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			data = append(data, DataAP{
				BSSID:      stmt.ColumnText(0),
				SSID:       stmt.ColumnText(1),
				Mode:       stmt.ColumnText(2),
				Discovered: stmt.ColumnText(3),
				Channel:    stmt.ColumnInt64(4),
				Frequency:  stmt.ColumnInt64(5),
				RSSI:       stmt.ColumnFloat(6),
				Latitude:   stmt.ColumnFloat(7),
				Longitude:  stmt.ColumnFloat(8),
				Altitude:   stmt.ColumnFloat(9),
				Accuracy:   stmt.ColumnFloat(10),
				Device:     stmt.ColumnText(11),
				Wigle:      stmt.ColumnBool(12),
				BeaconDB:   stmt.ColumnBool(13),
				DWPA:       stmt.ColumnBool(14),
			})
			return nil
		},
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - select ap exec failed")
	}

	return data
}
