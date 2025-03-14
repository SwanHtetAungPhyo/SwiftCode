package routes

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/SwanHtetAungPhyo/swifcode/internal/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.uber.org/zap"
)

func SetUpRoute(app *fiber.App, handlers *handler.SwiftCodeHandlers) {
	logging.Logger.Info("SetUpRoute", zap.String("app", app.MountPath()))
	apiVersioning := app.Group("/v1/swift-codes")
	apiVersioning.Get("/", handlers.Get)
	apiVersioning.Get("/country/:countryISO2code", handlers.GetWithISO2)
	apiVersioning.Post("/", handlers.Create)
	apiVersioning.Delete("/", handlers.Delete)

	app.Get("/metrics", monitor.New())
}
