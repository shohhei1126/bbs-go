package logger

import (
	"github.com/Sirupsen/logrus"
	"os"
)

type Conf struct {
	FilePath string
	LogLevel string
}

func NewLogger(conf Conf) (*logrus.Logger, error) {
	log := logrus.New()
	if conf.FilePath != "" {
		f, err := os.OpenFile(conf.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		log.Out = f
	}
	var err error
	log.Level, err = logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		return nil, err
	}
	return log, nil
}
