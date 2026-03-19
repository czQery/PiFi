package cmd

import (
	"context"
	"errors"
)

func SetMonitor(ctx context.Context, iface string, region string, auto bool) error {
	err := SetInterfaceMode(ctx, iface, "monitor")
	if err != nil {
		return errors.New("set interface mode: " + err.Error())
	}

	if err = SetBettercap(ctx, "set wifi.region "+region); err != nil {
		return errors.New("set region: " + err.Error())
	}
	if err = SetBettercapMonitor(ctx, iface); err != nil {
		return errors.New("set recon: " + err.Error())
	}

	MonitorAuto = auto
	return nil
}
