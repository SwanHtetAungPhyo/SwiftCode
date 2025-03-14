package logging

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger
var err error

func Init() {
	Logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
}

func SyncLogger() {
	if Logger != nil {
		defer func(Logger *zap.Logger) {
			err := Logger.Sync()
			if err != nil {
				panic(err)
			}
		}(Logger)
	}
}
