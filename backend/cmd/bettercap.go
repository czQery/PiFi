package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func InitBettercap() {
	ctx := context.Background()
	init := true
	for {
		logrus.Info("cmd - starting bettercap")

		eval := "set api.rest.address 127.0.0.1;set api.rest.port 8081;api.rest on;set ticker.period 60"
		cmd := exec.Command("bettercap", "-no-history", "-no-colors", "-eval", eval)

		stdout, _ := cmd.StdoutPipe()
		errStart := cmd.Start()

		if !init {
			go func() {
				// wait few seconds so bettercap can start
				time.Sleep(5 * time.Second)

				ifaceConfig := hp.ConfigGetInterfaceList()
				for iface, item := range ifaceConfig {
					if item["ready"] != true || item["mode"] != "monitor" {
						continue
					}

					logrus.WithFields(logrus.Fields{
						"iface": iface,
					}).Info("cmd - bettercap recovering monitor")

					err := SetBettercapMonitor(iface)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"iface": iface,
							"err":   err,
						}).Error("cmd - bettercap monitor recovery failed")
					}
					return
				}
			}()
		}

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
			case "wifi.client.probe":
				bssid, ssid, rssi := hp.ParseProbe(strings.TrimSpace(matches[4]))
				if bssid == "" {
					break
				}

				logrus.WithFields(logrus.Fields{
					"bssid": bssid,
					"ssid":  ssid,
					"rssi":  rssi,
				}).Debug("cmd - bettercap probe detected")
			case "wifi.ap.new":
				bssid, ssid, rssi := hp.ParseAP(strings.TrimSpace(matches[4]))
				if bssid == "" {
					break
				}

				logrus.WithFields(logrus.Fields{
					"bssid": bssid,
					"ssid":  ssid,
					"rssi":  rssi,
				}).Debug("cmd - bettercap ap detected")

				details, detailsErr := GetBettercapAP(bssid)
				if detailsErr != nil {
					logrus.WithFields(logrus.Fields{
						"bssid": bssid,
						"err":   detailsErr,
					}).Warn("cmd - bettercap ap get details failed")
				}

				auth := details.Get("authentication").String()
				if auth == "UNK" {
					auth = "PSK"
				}

				mode := fmt.Sprintf("[%s-%s-%s]", details.Get("encryption").String(), auth, details.Get("cipher").String())

				if details.Get("wps.State").Exists() {
					mode += "[WPS]"
				}

				mode += "[ESS]"
				mode = strings.ReplaceAll(mode, "--]", "]")
				mode = strings.ReplaceAll(mode, "-]", "]")

				var (
					lat float64 = 0
					lon float64 = 0
					alt float64 = 0
					acc float64 = 0
				)

				// use real gps data if they are less or equal 5 seconds old
				if time.Now().Unix()-GPS.Time <= 5 {
					lat = GPS.Lat
					lon = GPS.Lon
					alt = GPS.Alt
					acc = GPS.Acc
				}

				db.InsertAP(ctx, bssid, ssid, mode, time.Now().Format(time.DateTime), details.Get("channel").Int(), details.Get("frequency").Int(), rssi, lat, lon, alt, acc, "WIFI")
			}
		}

		init = false
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

func SetBettercapMonitor(iface string) error {
	_ = SetBettercap("ticker off")
	return SetBettercap("set wifi.interface " + iface + ";wifi.recon on;set ticker.commands 'wifi.recon on';ticker on")
}

func DisableBettercapMonitor() error {
	_ = SetBettercap("ticker off")
	return SetBettercap("set wifi.interface null;wifi.recon off")
}

func GetBettercapAP(bssid string) (gjson.Result, error) {
	rsp, err := req.Get(BC + "/api/session/wifi")
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.Parse(rsp.String()).Get(`aps.#(mac="` + bssid + `")`), nil
}
