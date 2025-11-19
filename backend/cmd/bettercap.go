package cmd

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func InitBettercap() {
	for {
		logrus.Info("cmd - starting bettercap")

		eval := "set api.rest.address 127.0.0.1;set api.rest.port 8081;api.rest on;set ticker.period 60"
		cmd := exec.Command("bettercap", "-no-history", "-no-colors", "-eval", eval)

		stdout, _ := cmd.StdoutPipe()
		errStart := cmd.Start()

		scanner := bufio.NewScanner(stdout)
		scannerRegex := regexp.MustCompile(`^\[(?P<time>[^]]+)]\s+\[(?P<module>[^]]+)]\s+(?:\[(?P<level>[^]]+)]\s+)?(?P<msg>.*)$`)
		for scanner.Scan() {
			line := scanner.Text()

			matches := scannerRegex.FindStringSubmatch(line)
			if matches == nil || len(matches) < 4+1 {
				continue
			}

			if strings.Contains(matches[2], "wifi") {
				logrus.WithFields(logrus.Fields{
					"module": strings.TrimSpace(matches[2]),
					"msg":    matches[4],
				}).Debug("cmd - bettercap out")
			}

			//fmt.Printf("Time: %s | Module: %-12s | Level: %-5s | Msg: %s\n", matches[1], matches[2], matches[3], matches[4])
		}

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
