package cmd

import (
	"context"
	"errors"
	"strconv"
)

func SetMonitor(ctx context.Context, iface string, region string, wardrive, channelHop bool, channel int) error {

	// this will cause race condition when setting multiple times at same time, but i don't care at this moment
	IsBettercapPaused = true
	defer func() {
		IsBettercapPaused = false
	}()

	err := SetInterfaceMode(ctx, iface, "monitor")
	if err != nil {
		return errors.New("set interface mode: " + err.Error())
	}

	if err = SetBettercap(ctx, "set wifi.region "+region, true); err != nil {
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
