package gps

import (
	"bufio"
	"github.com/czQery/PiFi/backend/hp"
	"github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"strings"
)

var COM serial.Port

func SerialListen() {
	var (
		read   byte
		buffer string
		err    error
	)

	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	COM, err = serial.Open(hp.Config.String("wigle.gps"), mode)
	if err != nil {
		list, _ := serial.GetPortsList()
		logrus.WithFields(logrus.Fields{
			"err":  err.Error(),
			"list": list,
		}).Error("gps - com failed to open")
		return
	}

	defer COM.Close()

	reader := bufio.NewReader(COM)

	for {
		read, err = reader.ReadByte()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err.Error(),
			}).Error("gps - read failed")
			return
		}

		buffer += string(read)

		if !strings.Contains(buffer, "\r\n") {
			continue
		}

		parts := strings.Split(buffer, "\r\n")
		for i, part := range parts {
			if i == len(part)-1 {
				buffer = part
				break
			}

			// ikd why the data is incomplete sometimes
			if !strings.Contains(part, "$") || len(part) < 3 {
				continue
			}

			logrus.WithFields(logrus.Fields{
				"data": part,
			}).Debug("gps - read")
		}
	}
}
