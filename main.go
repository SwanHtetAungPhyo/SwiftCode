package main

import (
	"github.com/SwanHtetAungPhyo/swifcode/cmd"
	"github.com/SwanHtetAungPhyo/swifcode/internal/config"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/SwanHtetAungPhyo/swifcode/pkg/logging"
)

// @title           SwiftCode API
// @version         1.0
// @description     This is the api to empower the bank system

// @contact.name   API Support
// @contact.email  swanhtetaungp@gmail.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080

func main() {
	// ENV LOADING
	configuration := config.LoadConfig()

	logger.Init()
	appLogger := logger.GetLogger()
	appLogger.Info("Logger initialized successfully")

	repo.Init(appLogger, configuration)
	db := repo.GetDBInstance()
	bankProcessor := services.NewBankProcessor(db, appLogger)
	bankProcessor.ProcessData(configuration.FilePath)
	appLogger.Info("Database initialized successfully")

	appLogger.Info("Starting SwiftCode API server on port :8080")
	if configuration.Mode == "development" {
		cmd.Start(configuration.PORT, appLogger, db, true)
	} else {
		cmd.Start(configuration.PORT, appLogger, db, false)
	}
}
