package main

import (
	"github.com/SwanHtetAungPhyo/swifcode/cmd"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/logging"
)

func main() {

	logging.Init()
	logger := logging.GetLogger()
	logger.Info("Logger initialized successfully")

	repo.Init(logger)
	logger.Info("Database initialized successfully")

	services.DataProcessing(logger)
	logger.Info("Starting SwiftCode API server on port :8080")
	cmd.Start(":8080", logger)
}
