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

bettercap:

- set api.rest.port 8088
- set api.rest.address 127.0.0.1
- set api.rest.username user
- set api.rest.password pass
- api.rest on

- set http.server.port 88
- set http.server.path /usr/share/bettercap/ui
- http.server on

- set wifi.interface wlxe4beed4c52e7
- wifi.recon on

- sudo bettercap -eval "set api.rest.address 0.0.0.0;api.rest on;set wifi.interface wlxe4beed4c52e7;wifi.recon on"

# Setup

NetworkManager: (this isn't really needed idk why i put it here)

- in /etc/NetworkManager/NetworkManager.conf
- `[main].dns=dnsmasq`

# GPS

- https://www.davidpilling.com/wiki/index.php/GPS
- https://austinsnerdythings.com/2025/02/14/revisiting-microsecond-accurate-ntp-for-raspberry-pi-with-gps-pps-in-2025/
