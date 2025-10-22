<p align="center" style="text-align: center">
  <img src="https://github.com/czQery/PiFi/blob/main/assets/dash.png?raw=true" alt="Banner">
  <br>
  <a href="https://goreportcard.com/report/github.com/czQery/PiFi/backend">
      <img src="https://goreportcard.com/badge/github.com/czQery/PiFi/backend" alt="report"/>
  </a>
  <a href="https://github.com/czQery/PiFi/actions" style="text-decoration: none">
    <img src="https://img.shields.io/github/actions/workflow/status/czQery/PiFi/release.yml" alt="build"/>
  </a>
  <a href="https://github.com/czQery/PiFi/releases/latest" style="text-decoration: none">
    <img src="https://img.shields.io/github/v/release/czQery/PiFi?include_prereleases" alt="release"/>
  </a>
  <br>
</p>

# ⚠️ In Early development!

# Setup

### 1. Install Rasbian

- Download [Raspberry Pi OS Lite](https://www.raspberrypi.com/software/operating-systems/)
- Flash using [balenaEtcher](https://etcher.balena.io/)
- Create `ssh` file in boot partition
- Create `userconf` file with data: `username:hashed-password`
  - Use this command to generate hashed-password
    ```bash
    echo 'password' | openssl passwd -6 -stdin
    ```

### 2. Update system

- Edit hostname in `/etc/hosts` & `/etc/hostname`
  ```bash
  sudo apt update
  sudo apt upgrade
  sudo reboot
  ```

- (Optional) Switch from `wpa_supplicant` to `iwd` if encountering wifi issues

### 3. Install PiFi

- Download [latest release](https://github.com/czQery/PiFi/releases)
  ```bash
  curl -L https://github.com/czQery/PiFi/releases/download/v0.0.4/release-arm.tar.gz | tar -xzv
  ```

### 4. Create service

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
- Reboot
  ```bash
  sudo reboot
  ```
