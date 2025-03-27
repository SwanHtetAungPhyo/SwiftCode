package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func Init() {
	once.Do(func() {
		logger = logrus.New()
		logger.SetOutput(os.Stdout)
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableHTMLEscape: true,
			PrettyPrint:       true,
		})

	})
}

func GetLogger() *logrus.Logger {
	if logger == nil {
		panic("Logger is not initialized. Call Init() before using GetLogger()")
	}
	return logger
}
