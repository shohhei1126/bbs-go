package log

import (
	"github.com/Sirupsen/logrus"
	"os"
)

type Conf struct {
	LogFile  string
	LogLevel string
}

var Logger = logrus.New()

func Init(conf Conf) error {
	if conf.LogFile != "" {
		f, err := os.OpenFile(conf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		Logger.Out = f
	}

	if conf.LogLevel != "" {
		var err error
		Logger.Level, err = logrus.ParseLevel(conf.LogLevel)
		if err != nil {
			return err
		}
	}
	return nil
}
