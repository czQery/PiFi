package main

import (
	"time"

	"github.com/czQery/PiFi/backend/api"
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/sirupsen/logrus"
)

func Timer() {
	for range time.Tick(3 * time.Second) {
		api.DeauthTargets.Range(func(key, _ interface{}) bool {
			err := cmd.SetBettercap("wifi.deauth " + key.(string))
			if err != nil {
				api.DeauthTargets.Delete(key)
				logrus.WithFields(logrus.Fields{
					"target": key.(string),
					"err":    err.Error(),
				}).Error("cmd - deauth failed")
			}
			return true
		})
	}
}
