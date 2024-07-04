package hp

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const LogFileName = "log.txt"

var LogFile *os.File

type LogFormatterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

func (hook *LogFormatterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *LogFormatterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
