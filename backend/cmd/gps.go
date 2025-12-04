package cmd

import (
	"bufio"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type GPSData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Alt float64 `json:"alt"`

	Mode int64 `json:"mode"`
	Time int64 `json:"time"`
}

var GPS GPSData

func InitGPS() {
	for {
		logrus.Info("cmd - starting gps")
		cmd := exec.Command("gpspipe", "-w")

		stdout, _ := cmd.StdoutPipe()
		errStart := cmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			data := gjson.Parse(line)

			switch data.Get("class").String() {
			case "TPV":
				if data.Get("mode").Int() == 0 { // ignore when not locked
					break
				}

				GPS = GPSData{
					Lat: data.Get("lat").Float(),
					Lon: data.Get("lon").Float(),
					Alt: data.Get("alt").Float(),

					Mode: data.Get("mode").Int(),
					Time: time.Now().Unix(),
				}

				/*logrus.WithFields(logrus.Fields{
					"lat": data.Get("lat").Float(),
					"lon": data.Get("lon").Float(),
				}).Debug("cmd - gps out")*/
			}
		}

		errWait := cmd.Wait()
		logrus.WithFields(logrus.Fields{
			"errStart": errStart,
			"errWait":  errWait,
		}).Error("cmd - gps exited!")

		time.Sleep(5 * time.Second)
	}
}
