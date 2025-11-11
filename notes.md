# Commands

device list:

- nmcli -t device status

hotspot setup:

- sudo nmcli con add type wifi ifname wlan0 con-name PiFi-hotspot autoconnect yes ssid kvn
- sudo nmcli con modify PiFi-hotspot 802-11-wireless.mode ap 802-11-wireless.band bg ipv4.method shared

password setting:

- sudo nmcli con modify PiFi-hotspot wifi-sec.key-mgmt wpa-psk wifi-sec.psk "12345678" wifi-sec.pmf disable

channel setting:

- sudo nmcli con modify PiFi-hotspot 802-11-wireless.channel 5

dns setting:

- /etc/NetworkManager/dnsmasq-shared.d/PiFi.conf
- `address=/#/10.42.0.1`
- sudo systemctl restart NetworkManager

enable hotspot:

- sudo nmcli con up PiFi-hotspot

client setup:

- sudo nmcli device wifi rescan ifname wlan0
- sudo nmcli device wifi connect ssid password password ifname wlan0 hidden yes name PiFi-client

monitor:

- sudo ifconfig wlan0 down
- sudo iwconfig wlan0 mode monitor
- sudo airodump-ng wlan1 --manufacturer --wps

# Setup

NetworkManager

- in /etc/NetworkManager/NetworkManager.con
- `[main].dns=dnsmasq`
