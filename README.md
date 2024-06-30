# PiFi
Raspberry pi zero wifi pentesting kit


# Commands

device list:
- nmcli -t device status

hotspot setup:
- sudo nmcli con add type wifi ifname wlan0 con-name PiFi autoconnect yes ssid kvn
- sudo nmcli con modify PiFi 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared

password setting:
- sudo nmcli con modify PiFi wifi-sec.key-mgmt wpa-psk
- sudo nmcli con modify PiFi wifi-sec.psk "12345678"

channel setting:
- sudo nmcli con modify PiFi 802-11-wireless.channel 5

- enable hotspot:
sudo nmcli con up PiFi