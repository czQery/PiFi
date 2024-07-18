package cmd

import (
	"github.com/czQery/PiFi/backend/hp"
	"os"
)

func SetDNSPortal() error {
	data := []byte("address=/#/" + hp.Config.String("main.gateway") + "\n")
	err := os.WriteFile("/etc/NetworkManager/dnsmasq-shared.d/"+Con+".conf", data, 0644)
	return err
}

func DisableDNSPortal() error {
	data := []byte("")
	err := os.WriteFile("/etc/NetworkManager/dnsmasq-shared.d/"+Con+".conf", data, 0644)
	return err
}
