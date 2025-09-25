package main

import (
	"time"

	"github.com/czQery/PiFi/backend/api"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/sirupsen/logrus"
)

func ticker() {

	var err error

	for range time.Tick(5 * time.Second) {
		api.StatsCPU, err = cpu.Get()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Warn("ticker - cpu stats failed")
		}
	}
}
