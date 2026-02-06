package hp

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const LogFileName = "log.txt"

var LogFile *os.File

type LogFormatterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
	Broadcast func([]byte)
}

func (hook *LogFormatterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}

	if hook.Broadcast != nil {
		hook.Broadcast(line)
	}

	_, err = hook.Writer.Write(line)
	return err
}

func (hook *LogFormatterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
