package cmd

import (
	"bufio"
	"math"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type GPSData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Alt int64   `json:"alt"`
	Acc float64 `json:"acc"`

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
				GPS.Mode = data.Get("mode").Int()
				GPS.Time = data.Get("time").Int()

				// by my findings it looks like eph, epx or eph are only returned in mode 3, so manual calculation of accuracy from SKY class will probably be needed
				if data.Get("eph").Exists() {
					GPS.Acc = data.Get("eph").Float()
				} else if data.Get("epx").Exists() && data.Get("epy").Exists() {
					epx := data.Get("epx").Float()
					epy := data.Get("epy").Float()
					GPS.Acc = math.Sqrt((epx * epx) + (epy * epy))
				}

				switch data.Get("mode").Int() {
				case 3:
					GPS.Alt = data.Get("alt").Int()
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
