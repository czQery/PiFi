package cmd

import (
	"os/exec"
	"time"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func InitBettercap() {
	for {
		logrus.Info("cmd - starting bettercap")

		eval := "set api.rest.address 127.0.0.1;api.rest.port 8081;api.rest on"
		cmd := exec.Command("bettercap", "-no-history", "-eval", eval)

		errStart := cmd.Start()
		errWait := cmd.Wait()

		logrus.WithFields(logrus.Fields{
			"errStart": errStart,
			"errWait":  errWait,
		}).Error("cmd - bettercap exited!")

		time.Sleep(5 * time.Second)
	}
}

func SetBettercap(cmd string) error {
	body := "{\"cmd\":\"" + cmd + "\"}"
	_, err := req.SetBodyJsonString(body).Post(BC + "/api/session")
	if err != nil {
		return err
	}

	return nil
}
