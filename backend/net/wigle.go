package net

import (
	"errors"
	"strconv"
	"time"

	"github.com/czQery/PiFi/backend/cmd"
	"github.com/czQery/PiFi/backend/db"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func WigleGet() string {
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

	for _, ap := range db.SelectAP() {
		if ap.Wigle { // skip if already uploaded
			continue
		}

		payload += ap.BSSID + ","
		payload += ap.SSID + ","
		payload += ap.Mode + ","
		payload += ap.Discovered + ","
		payload += strconv.FormatInt(ap.Channel, 10) + ","
		payload += strconv.FormatInt(ap.Frequency, 10) + ","
		payload += strconv.FormatFloat(ap.RSSI, 'f', -1, 64) + ","
		payload += strconv.FormatFloat(ap.Latitude, 'f', -1, 64) + ","
		payload += strconv.FormatFloat(ap.Longitude, 'f', -1, 64) + ","
		payload += strconv.FormatInt(ap.Altitude, 10) + ","
		payload += strconv.FormatFloat(ap.Accuracy, 'f', -1, 64) + ","
		payload += "," // RCOIs
		payload += "," // MfgrId
		payload += ap.Device + "\n"
	}

	return payload
}

func WigleUpload() error {
	name := "PiFi_" + time.Now().Format("20060102150405") + ".csv"
	token := "Basic " + hp.Config.Get("net.wigle").(string)

	logrus.WithFields(logrus.Fields{
		"name":  name,
		"token": token,
	}).Debug("net - wigle upload")

	up, err := req.SetHeader("Authorization", token).SetFileBytes("file", name, []byte(WigleGet())).Post("https://api.wigle.net/api/v2/file/upload")
	if err != nil {
		return err
	}

	upData := gjson.Parse(up.String())

	if up.StatusCode != 200 || !upData.Get("success").Bool() {
		return errors.New("upload failed: " + upData.Get("message").String())
	}

	return nil
}

/*	UPDATE your_table
		SET your_bool_column = 1
		WHERE id IN (
	    	SELECT value FROM json_each(?)
		);
*/
