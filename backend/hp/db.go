package hp

import (
	"strings"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"

	"github.com/sirupsen/logrus"
)

var DB *sqlite.Conn

const dbCreate = `
	CREATE TABLE IF NOT EXISTS ap (
		bssid TEXT NOT NULL PRIMARY KEY,
		ssid TEXT NOT NULL,
		mode TEXT NOT NULL,
		discovered TEXT NOT NULL,
		channel INTEGER NOT NULL,
		rssi REAL NOT NULL,
		latitude REAL NOT NULL,
		longitude REAL NOT NULL,
		altitude INTEGER NOT NULL,
		accuracy INTEGER NOT NULL,
		device TEXT NOT NULL
	);
`

func DBLoad() {
	var err error
	DB, err = sqlite.OpenConn("./data.db", sqlite.OpenReadWrite|sqlite.OpenCreate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - open failed")
	}

	err = sqlitex.Execute(DB, strings.TrimSpace(dbCreate), &sqlitex.ExecOptions{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - exec failed")
	}

	// TEST
	// DBInsertAP("a8:80:55:42:e8:e2", "ap", "wifi", "now", 1, 5, 10, 20, 400, 5, "wifi")
}

func DBInsertAP(bssid, ssid, mode, discovered string, channel int, rssi, latitude, longitude float64, altitude, accuracy int, device string) {
	err := sqlitex.Execute(DB, "INSERT INTO ap (bssid, ssid, mode, discovered, channel, rssi, latitude, longitude, altitude, accuracy, device) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", &sqlitex.ExecOptions{
		Args: []interface{}{
			bssid,
			ssid,
			mode,
			discovered,
			channel,
			rssi,
			latitude,
			longitude,
			altitude,
			accuracy,
			device,
		},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("db - insert_db exec failed")
	}
}
