package db

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

var Pool *sqlitex.Pool

func Load() {
	var err error

	Pool, err = sqlitex.NewPool("./data.db", sqlitex.PoolOptions{
		PoolSize: 10,
		Flags:    sqlite.OpenReadWrite | sqlite.OpenCreate,
	})

	conn, err := Pool.Take(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - connection failed")
	}
	defer Pool.Put(conn)

	err = sqlitex.Execute(conn, strings.TrimSpace(dbCreate), &sqlitex.ExecOptions{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Panic("db - exec failed")
	}

	logrus.Info("db - successfully loaded")
}
