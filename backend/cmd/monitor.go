package cmd

import (
	"context"
	"errors"
	"strconv"
)

func SetMonitor(ctx context.Context, iface string, region string, wardrive, channelHop bool, channel int) error {
	err := SetInterfaceMode(ctx, iface, "monitor")
	if err != nil {
		return errors.New("set interface mode: " + err.Error())
	}

	if err = SetBettercap(ctx, "set wifi.region "+region); err != nil {
		return errors.New("set region: " + err.Error())
	}

	var channelString string
	if channelHop {
		channelString = "clear"
	} else {
		channelString = strconv.Itoa(channel)
	}

	if err = SetBettercapMonitor(ctx, iface, channelString); err != nil {
		return errors.New("set recon: " + err.Error())
	}

	MonitorWardrive = wardrive
	MonitorChannelHop = channelHop
	MonitorChannel = channel
	return nil
}
