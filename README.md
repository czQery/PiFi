<div align="center" style="text-align: center">
  <img src="https://github.com/czQery/PiFi/blob/main/assets/banner.png?raw=true" alt="Banner">
  <br>

[![Badge Report]][Report]
[![Badge Build]][Build]
[![Badge Release]][Release]

</div>

# Setup

### 1. Install Raspbian

- Flash using [rpi-imager](https://www.raspberrypi.com/software/)
- Set username, password, hostname etc.
- Enable ssh

### 2. Setup system

- Install updates
  ```bash
  sudo apt update
  sudo apt upgrade
  sudo rpi-update
  sudo reboot
  ```
- Use `raspi-config` to set wlan country
- Use `raspi-config` to enable predictable interface names

### 3. Install Dependencies

```bash
sudo apt install bettercap bettercap-caplets bettercap-ui
sudo systemctl stop bettercap
sudo systemctl disable bettercap
```

### 4. Install PiFi

- Download [latest release](https://github.com/czQery/PiFi/releases)
  ```bash
  curl -L https://github.com/czQery/PiFi/releases/download/v0.0.4/release-arm.tar.gz | tar -xzv
  ```

### 5. Create service

- Edit the example `pifi.service` file in this repo and put it in `/etc/systemd/system/`
- Reload systemd daemon
  ```bash
  sudo systemctl daemon-reload
  ```
- Enable service
  ```bash
  sudo systemctl enable pifi.service
  ```
- Start service
  ```bash
  sudo systemctl start pifi.service
  ```

# GPS

- Enable serial interface
  ```bash
  sudo raspi-config
  ```
- Select -> Interfacing Options
- Select -> Serial
- Login shell -> No
- Hardware Serial -> Yes
  ```bash
  sudo apt install pps-tools gpsd gpsd-clients chrony
  sudo systemctl stop gpsd
  sudo systemctl stop gpsd.socket

  # Edit these settings to match your gps module
  sudo ubxtool -S 460800 -s 38400 -f /dev/serial0
  sudo ubxtool -s 460800 -e GPS -e GALILEO -e GLONASS -e BEIDOU -e SBAS -e QZSS -f /dev/serial0
  sudo ubxtool -s 460800 -z CFG-RATE-MEAS,100 -f /dev/serial0
  sudo ubxtool -s 460800 -z CFG-NAVSPG-DYNMODEL,4 -f /dev/serial0
  sudo ubxtool -s 460800 -p SAVE -f /dev/serial0
  ```
- Edit `/etc/default/gpsd`
  ```bash
  DEVICES="/dev/serial0"
  GPSD_OPTIONS="-n -s 460800"
  USBAUTO="true"
  ```
- Apply
  ```bash
  sudo systemctl enable gpsd
  sudo reboot
  ```


<!----------------------------------------------------------------------------->

[Report]: https://goreportcard.com/report/github.com/czQery/PiFi/backend
[Badge Report]: https://goreportcard.com/badge/github.com/czQery/PiFi/backend

[Build]: https://github.com/czQery/PiFi/action
[Badge Build]: https://img.shields.io/github/actions/workflow/status/czQery/PiFi/release.yml

[Release]: https://github.com/czQery/PiFi/releases/latest
[Badge Release]: https://img.shields.io/github/v/release/czQery/PiFi

<!----------------------------------------------------------------------------->