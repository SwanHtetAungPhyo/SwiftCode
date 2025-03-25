package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

func Init() {
	logger = logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:             true,
		PadLevelText:              true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		QuoteEmptyFields:          true,
		TimestampFormat:           "2006-01-02 15:04:05",
	})
}

func GetLogger() *logrus.Logger {
	return logger
}
