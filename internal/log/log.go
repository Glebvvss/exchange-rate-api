package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(os.Getenv("LOG_PATH"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Error log file was not oppened and uses stderr now")
	}
}

func Error(err error) {
	log.Error(err)
}
