package main

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/utils"
	"github.com/SwanHtetAungPhyo/swifcode/internal/repo"
	"github.com/SwanHtetAungPhyo/swifcode/internal/routes"
	"github.com/SwanHtetAungPhyo/swifcode/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	portNumber := os.Getenv("PORT")
	dataPath := os.Getenv("DATA_PATH")
	parsedData := utils.Parse(dataPath)

	logging.Init()
	defer logging.SyncLogger()
	app := apiInit()
	repo.Init()
	if repo.DbInstance == nil {
		log.Fatal("Failed to get DB instance")
	}

	go utils.InsertData(parsedData)

	serviceInstanceForInject := services.NewSwiftCodeService(repo.DbInstance)
	handlerInstance := handler.NewSwiftCodeHandlers(*serviceInstanceForInject)

	routes.SetUpRoute(app, handlerInstance)

	go func() {
		logging.Logger.Info("Starting server", zap.String("address", "http://localhost:"+portNumber))
		err := app.Listen(":" + portNumber)
		if err != nil {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-signalChan

	logging.Logger.Info("Shutting down...")
	if err := app.Shutdown(); err != nil {
		logging.Logger.Error("Error shutting down", zap.Error(err))
	}
}

// Configuration fot the best and proper speed
func apiInit() *fiber.App {
	appConfig := fiber.New(fiber.Config{
		ServerHeader: "SWIFT_CODE",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Prefork:      true,
	})
	// rate limiting
	appConfig.Use(limiter.Config{
		Max:        30,
		Expiration: 5 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "too many requests",
			})
		},
	})
	//
	appConfig.Use(compress.New(
		compress.Config{
			Level: compress.LevelBestSpeed,
		}))
	
	return appConfig
}
