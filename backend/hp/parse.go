package hp

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseNewAP(input string) (bssid, ssid string, rssi float64) {
	regex := regexp.MustCompile(`^wifi access point (?P<ssid>.+) \((?P<rssi>-[0-9]+ dBm)\) detected as (?P<bssid>[0-9a-fA-F:]{17})\.$`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil || len(matches) < 4 {
		return
	}

	bssid = strings.TrimSpace(matches[3])
	ssid = strings.TrimSpace(matches[1])
	rssi, _ = strconv.ParseFloat(strings.TrimSpace(strings.Replace(matches[2], "dBm", "", 1)), 64)

	return
}
