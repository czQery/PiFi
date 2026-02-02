package net

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func WigleGet() (string, []string) {
	payload := "WigleWifi-1.6"
	payload += ",appRelease=" + hp.Build
	payload += ",model=" + cmd.Con
	payload += ",release=" + hp.Build
	payload += ",device=" + cmd.Con
	payload += ",display="
	payload += ",board=RaspberryPi"
	payload += ",brand=PiFi"
	payload += ",star=Sol"
	payload += ",body=4"
	payload += ",subBody=1\n"

	payload += "MAC,SSID,AuthMode,FirstSeen,Channel,Frequency,RSSI,CurrentLatitude,CurrentLongitude,AltitudeMeters,AccuracyMeters,RCOIs,MfgrId,Type\n"

	var list []string

	for _, ap := range db.SelectAP() {
		if ap.Wigle { // skip if already uploaded
			continue
		}

		if ap.Accuracy == 0 && ap.Longitude == 0 && ap.Latitude == 0 { // skip data without geolocation
			continue
		}

		payload += ap.BSSID + ","
		payload += "\"" + ap.SSID + "\","
		payload += ap.Mode + ","
		payload += ap.Discovered + ","
		payload += strconv.FormatInt(ap.Channel, 10) + ","
		payload += strconv.FormatInt(ap.Frequency, 10) + ","
		payload += strconv.FormatFloat(ap.RSSI, 'f', -1, 64) + ","
		payload += strconv.FormatFloat(ap.Latitude, 'f', -1, 64) + ","
		payload += strconv.FormatFloat(ap.Longitude, 'f', -1, 64) + ","
		payload += strconv.FormatInt(int64(math.Round(ap.Altitude)), 10) + "," // only gps & wigle mismatch, wigle wants int, gps gives float
		payload += strconv.FormatFloat(ap.Accuracy, 'f', -1, 64) + ","
		payload += "," // RCOIs
		payload += "," // MfgrId
		payload += ap.Device + "\n"

		list = append(list, ap.BSSID)
	}

	return payload, list
}

func WigleUpload() error {
	name := "PiFi_" + time.Now().Format("20060102150405") + ".csv"
	token := ""

	if hp.Config.Get("net.wigle") != nil {
		token = "Basic " + hp.Config.Get("net.wigle").(string)
	}

	logrus.WithFields(logrus.Fields{
		"name":  name,
		"token": token,
	}).Debug("net - wigle upload")

	payload, list := WigleGet()

	up, err := req.SetHeader("Authorization", token).SetFormData(map[string]string{"donate": "on"}).SetFileBytes("file", name, []byte(payload)).Post("https://api.wigle.net/api/v2/file/upload")
	if err != nil {
		return err
	}

	upData := gjson.Parse(up.String())

	if up.StatusCode != 200 || !upData.Get("success").Bool() {
		return errors.New("upload failed: " + upData.Get("message").String())
	}

	var transids []string
	for _, id := range upData.Get("results.transids.#.transId").Array() {
		transids = append(transids, id.String())
	}

	logrus.WithFields(logrus.Fields{
		"name":     name,
		"observer": upData.Get("observer").String(),
		"warning":  upData.Get("warning").String(),
		"transids": strings.Join(transids, ","),
	}).Info("net - wigle successfully uploaded")

	WigleMark(true, list)

	return nil
}

func WigleMark(value bool, list []string) {
	db.UpdateAPNet("wigle", value, list)
}
