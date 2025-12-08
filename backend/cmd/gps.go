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

		GPS = GPSData{}
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			data := gjson.Parse(line)

			switch data.Get("class").String() {
			case "TPV":
				/*logrus.WithFields(logrus.Fields{
					"lat": data.Get("lat").Float(),
					"lon": data.Get("lon").Float(),
				}).Debug("cmd - gps out")*/

				GPS.Mode = data.Get("mode").Int()
				GPS.Time = data.Get("time").Int()

				switch data.Get("mode").Int() {
				case 3:
					GPS.Alt = data.Get("alt").Float()
					fallthrough
				case 2:
					GPS.Lat = data.Get("lat").Float()
					GPS.Lon = data.Get("lon").Float()
				}
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
