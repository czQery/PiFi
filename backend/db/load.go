package db

import (
	"strings"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

var Conn *sqlite.Conn

func Load() {
	var err error
	Conn, err = sqlite.OpenConn("./data.db", sqlite.OpenReadWrite|sqlite.OpenCreate)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - open failed")
	}

	err = sqlitex.Execute(Conn, strings.TrimSpace(dbCreate), &sqlitex.ExecOptions{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - exec failed")
	}

	// TEST
	// DBInsertAP("a8:80:55:42:e8:e2", "ap", "wifi", "now", 1, 5, 10, 20, 400, 5, "wifi")
}
