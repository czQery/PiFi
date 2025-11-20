package cmd

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/czQery/PiFi/backend/hp"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

var RunningBettercap bool

func InitBettercap() {
	for {
		logrus.Info("cmd - starting bettercap")

		eval := "set api.rest.address 127.0.0.1;set api.rest.port 8081;api.rest on;set ticker.period 60"
		cmd := exec.Command("bettercap", "-no-history", "-no-colors", "-eval", eval)

		stdout, _ := cmd.StdoutPipe()
		errStart := cmd.Start()

		RunningBettercap = true

		scanner := bufio.NewScanner(stdout)
		scannerRegex := regexp.MustCompile(`^\[(?P<time>[^]]+)]\s+\[(?P<module>[^]]+)]\s+(?:\[(?P<level>[^]]+)]\s+)?(?P<msg>.*)$`)
		for scanner.Scan() {
			line := scanner.Text()

			matches := scannerRegex.FindStringSubmatch(line)
			if matches == nil || len(matches) < 4+1 {
				continue
			}

			/*if strings.Contains(matches[2], "wifi") {
				logrus.WithFields(logrus.Fields{
					"module": strings.TrimSpace(matches[2]),
					"msg":    matches[4],
				}).Debug("cmd - bettercap out")
			}*/

			switch strings.TrimSpace(matches[2]) {
			case "wifi.ap.new":
				bssid, ssid, rssi := hp.ParseNewAP(strings.TrimSpace(matches[4]))
				if bssid == "" {
					break
				}

				logrus.WithFields(logrus.Fields{
					"bssid": bssid,
					"ssid":  ssid,
					"rssi":  rssi,
				}).Debug("cmd - bettercap new ap")

				hp.DBInsertAP(bssid, ssid, "wifi", time.Now().Format(time.DateTime), 1, rssi, 10, 20, 400, 5, "wifi")
			}
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
