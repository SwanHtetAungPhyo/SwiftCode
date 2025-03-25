package cmd

import (
	"context"
	"errors"
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/SwanHtetAungPhyo/swifcode/internal/middleware"
	"github.com/SwanHtetAungPhyo/swifcode/internal/model"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/routes"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(port string, log *logrus.Logger, dbInstance *gorm.DB, isDevelopmentMode bool) {
	log.Info("Initializing server components...")

	app := apiEngineInit()
	middleware.SetUp(app)

	handlers := dependencyInjection(log, dbInstance)
	routes.SetUpRoute(app, handlers, log)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Infof("Server is starting on port %s", port)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("Failed to start server")
		}
	}()

	if isDevelopmentMode {
		shutdown(server, log, dbInstance, true)
	} else {
		shutdown(server, log, dbInstance, false)
	}
}

func dependencyInjection(log *logrus.Logger, db *gorm.DB) *handler.SwiftCodeHandlers {
	log.Info("Initializing dependencies...")

	repoInst := repo.NewRepository(db, log)
	var repoMethods repo.RepositoryMethods = repoInst

	serviceInst := services.NewService(repoMethods, log)
	var serviceMethods services.ServiceMethods = serviceInst

	handlers := handler.NewSwiftCodeHandlers(serviceMethods, log)

	log.Info("Dependencies initialized successfully")
	return handlers
}

func shutdown(server *http.Server, log logrus.FieldLogger, db *gorm.DB, isDevelopment bool) {
	log.Info("Waiting for shutdown signal...")

	osChan := make(chan os.Signal, 1)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)
	<-osChan

	log.Warn("Shutdown signal received, initiating graceful shutdown...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("Error during HTTP server shutdown")
	}

	if isDevelopment {
		err := db.Migrator().DropTable(&model.BankDetails{}, &model.Town{}, &model.Country{})
		if err != nil {
			log.WithError(err).Fatal("Error purging database table")
			return
		}
	}

	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.WithError(err).Error("Failed to retrieve database connection")
		} else {
			if err := sqlDB.Close(); err != nil {
				log.WithError(err).Error("Failed to close database connection")
			} else {
				log.Info("Database connection closed successfully")
			}
		}
	}

	log.Info("Server shutdown completed successfully")
}
func apiEngineInit() *gin.Engine {
	log := logrus.New()
	log.Info("Setting up Gin engine...")
	router := gin.New()
	return router
}
