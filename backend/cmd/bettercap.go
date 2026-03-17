package cmd

import (
	"bufio"
	"cmp"
	"context"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"slices"
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

	eval := []string{
		"set api.rest.address 127.0.0.1",
		"set api.rest.port 8081",
		"set wifi.handshakes.aggregate false",
		"set wifi.handshakes.file ./cap",
		"set wifi.assoc.acquired true",
		"set wifi.assoc.open false",
		"set wifi.deauth.acquired true",
		"set wifi.deauth.open true",
		"set ticker.period 60",
		"api.rest on",
	}

	for {
		logrus.Info("cmd - starting bettercap")
		cmd := exec.CommandContext(ctx, "bettercap", "-no-history", "-no-colors", "-eval", strings.Join(eval, ";"))

		stdout, _ := cmd.StdoutPipe()
		errStart := cmd.Start()

		if !init {
			go func(ctx context.Context) {
				// wait few seconds so bettercap can start
				time.Sleep(5 * time.Second)

				ifaceConfig := hp.ConfigGetInterfaceList()
				for iface, item := range ifaceConfig {
					if item["ready"] != true || item["mode"] != "monitor" {
						continue
					}

					logrus.WithFields(logrus.Fields{
						"iface": iface,
					}).Info("cmd - recovering bettercap monitor")

					err := SetBettercapMonitor(ctx, iface)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"iface": iface,
							"err":   err.Error(),
						}).Error("cmd - bettercap monitor recovery failed")
					}
					return
				}
			}(ctx)
		}

		go func(stdout io.ReadCloser, ctx context.Context) {
			probes := make(map[string]struct{})
			handshakes := make(map[string]struct{})
			scanner := bufio.NewScanner(stdout)
			scannerRegex := regexp.MustCompile(`^\[(?P<time>[^]]+)]\s+\[(?P<module>[^]]+)]\s+(?:\[(?P<level>[^]]+)]\s+)?(?P<msg>.*)$`)

			for scanner.Scan() {
				line := scanner.Text()

				//fmt.Println("BTT", line)

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

				event := strings.TrimSpace(matches[2])
				data := strings.TrimSpace(matches[4])

				switch event {
				case "wifi.client.deauthentication":

					// TODO: Create parser for this event
					logrus.WithFields(logrus.Fields{
						"data": data,
					}).Debug("cmd - deauth detected")
				case "wifi.client.handshake":
					bssid, ssid, capType := hp.ParseHandshake(data)
					if bssid == "" {
						break
					}

					if _, ok := handshakes[bssid+ssid+capType]; ok {
						break
					}

					handshakes[bssid+ssid+capType] = struct{}{}
					db.UpdateAPHandshake(ctx, bssid, true)

					logrus.WithFields(logrus.Fields{
						"bssid": bssid,
						"ssid":  ssid,
						"type":  capType,
					}).Info("cmd - handshake captured")
				case "wifi.client.probe":
					bssid, ssid, _ := hp.ParseProbe(data)
					if bssid == "" {
						break
					}

					if _, ok := probes[bssid+ssid]; ok {
						break
					}

					probes[bssid+ssid] = struct{}{}
					logrus.WithFields(logrus.Fields{
						"bssid": bssid,
						"ssid":  ssid,
					}).Info("cmd - probe detected")
				case "wifi.ap.new":
					bssid, ssid, rssi := hp.ParseAP(data)
					if bssid == "" {
						break
					}

					details, detailsErr := GetBettercapAP(ctx, bssid)
					if detailsErr != nil {
						logrus.WithFields(logrus.Fields{
							"bssid": bssid,
							"err":   detailsErr.Error(),
						}).Warn("cmd - ap get details failed")
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
		}(stdout, ctx)

		init = false
		errWait := cmd.Wait()
		logrus.WithFields(logrus.Fields{
			"errStart": errStart,
			"errWait":  errWait,
		}).Error("cmd - bettercap exited!")

		time.Sleep(5 * time.Second)
	}
}

func SetBettercap(ctx context.Context, cmd string) error {
	body := "{\"cmd\":\"" + cmd + "\"}"
	_, err := req.NewClient().R().SetContext(ctx).SetBodyJsonString(body).Post(BC + "/api/session")
	if err != nil {
		return err
	}

	return nil
}

func SetBettercapMonitor(ctx context.Context, iface string) error {
	return SetBettercap(ctx, "set wifi.interface "+iface+";wifi.recon on")
}

func DisableBettercapMonitor(ctx context.Context) error {
	return SetBettercap(ctx, "set wifi.interface null;wifi.recon off")
}

func GetBettercapAP(ctx context.Context, bssid string) (gjson.Result, error) {
	rsp, err := req.NewClient().R().SetContext(ctx).Get(BC + "/api/session/wifi")
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.Parse(rsp.String()).Get(`aps.#(mac="` + bssid + `")`), nil
}

func GetBettercapAPs(ctx context.Context) ([]gjson.Result, error) {
	rsp, err := req.NewClient().R().SetContext(ctx).Get(BC + "/api/session/wifi")
	if err != nil {
		return []gjson.Result{}, err
	}

	result := gjson.Parse(rsp.String()).Get(`aps`).Array()

	slices.SortFunc(result, func(a, b gjson.Result) int {
		return cmp.Compare(a.Get("channel").Int(), b.Get("channel").Int())
	})

	return result, nil
}
