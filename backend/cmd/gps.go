package cmd

import (
	"bufio"
	"context"
	"math"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

type GPSData struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Alt float64 `json:"alt"`
	Acc float64 `json:"acc"`
	Spd float64 `json:"spd"`

	Mode int64 `json:"mode"`
	Time int64 `json:"time"`
}

var GPS GPSData

func InitGPS() {
	ctx := context.Background()

	for {
		logrus.Info("cmd - starting gps")
		cmd := exec.CommandContext(ctx, "gpspipe", "-w")

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

				timeParsed, _ := time.Parse(time.RFC3339, data.Get("time").String())
				GPS.Time = timeParsed.Unix()

				// eph reported sometimes even in mode 2, should be more accurate but i added epx & epy as fallback anyways
				if data.Get("eph").Exists() {
					GPS.Acc = data.Get("eph").Float()
				} else if data.Get("epx").Exists() && data.Get("epy").Exists() {
					epx := data.Get("epx").Float()
					epy := data.Get("epy").Float()
					GPS.Acc = math.Sqrt((epx * epx) + (epy * epy))
				}

				if data.Get("alt").Exists() {
					GPS.Alt = data.Get("alt").Float()
				}

				if data.Get("speed").Exists() {
					GPS.Spd = data.Get("speed").Float()
				}

				if data.Get("lat").Exists() && data.Get("lon").Exists() {
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
