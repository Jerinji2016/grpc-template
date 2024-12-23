package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

func InitLogger() {
	env := os.Getenv("ENVIRONMENT")
	// changed formatter to TextFormatter for readability
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
	)
	
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
	
	if env == "production" {
		file := getFile()

		log.SetOutput(
			&lumberjack.Logger{
				Filename:   file.Name(),
				MaxSize:    5,    // megabytes
				MaxBackups: 3,    // keep atmost 3 logs files
				MaxAge:     30,   // days to retain
				Compress:   true, // compress old log files
			},
		)
	}
}

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		createErr := os.MkdirAll(dir, os.ModePerm)
		if createErr != nil {
			log.Fatalf("Failed to create log directory: %v", createErr)
		}
	}
}

func getFile() *os.File {
	now := time.Now()

	logFileFormat := os.Getenv("LOG_FILE_FORMAT")
	if logFileFormat == "" {
		logFileFormat = "LOG_2006-01-02 15:04:05.log"
	}

	filepath := now.Format(logFileFormat)

	dir := os.Getenv("LOG_OUTPUT_DIR")
	if dir != "" {
		createDir(dir)
		filepath = fmt.Sprintf("%s/%s", dir, filepath)
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	return file
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

func FatalLog(format string, args ...interface{}) {
	//	write to file or something else
	log.Fatalf(format, args...)
}
