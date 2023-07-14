package log

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func New() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		log.SetLevel(logrus.InfoLevel)
	}
	return log
}
