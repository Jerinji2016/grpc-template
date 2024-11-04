package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

func InfoLog(format string, args ...interface{}) {
	//	write to file or something else
	log.Infof(format, args...)
}

func DebugLog(format string, args ...interface{}) {
	//	write to file or something else
	log.Debugf(format, args...)
}

func ErrorLog(format string, args ...interface{}) {
	//	write to file or something else
	log.Errorf(format, args...)
}

func WarnLog(format string, args ...interface{}) {
	//	write to file or something else
	log.Warnf(format, args...)
}