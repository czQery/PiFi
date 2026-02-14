package hp

import (
	"regexp"
	"strconv"
	"strings"
)

func ParseHandshake(input string) (bssid, ssid, capType string) {
	regex := regexp.MustCompile(`captured.*?->\s+(?P<ssid>.+?)\s+\((?P<bssid>[0-9a-fA-F:]{17})\)\s+(?P<type>.+?)\s+to`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil || len(matches) < 4 {
		return
	}

	ssid = strings.TrimSpace(matches[1])
	bssid = strings.TrimSpace(matches[2])
	capType = strings.TrimSpace(matches[3])
	return
}

func ParseProbe(input string) (bssid, ssid string, rssi float64) {
	regex := regexp.MustCompile(`station\s+(?P<bssid>[0-9a-fA-F:]{17}).*?SSID\s+(?P<ssid>.+)\s+\((?P<rssi>-[0-9]+)\s+dBm\)`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil || len(matches) < 4 {
		return
	}

	bssid = matches[1]
	ssid = strings.TrimSpace(matches[2])
	rssi, _ = strconv.ParseFloat(strings.TrimSpace(matches[3]), 64)
	return
}

func ParseAP(input string) (bssid, ssid string, rssi float64) {
	regex := regexp.MustCompile(`wifi access point\s+(?P<ssid>.+)\s+\((?P<rssi>-[0-9]+) dBm\)\s+detected as\s+(?P<bssid>[0-9a-fA-F:]{17})`)
	matches := regex.FindStringSubmatch(input)
	if matches == nil || len(matches) < 4 {
		return
	}

	bssid = strings.TrimSpace(matches[3])
	ssid = strings.TrimSpace(matches[1])
	rssi, _ = strconv.ParseFloat(strings.TrimSpace(matches[2]), 64)
	return
}
