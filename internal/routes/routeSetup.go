package routes

import (
	_ "github.com/SwanHtetAungPhyo/swifcode/docs"
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRoute(router *gin.Engine, handlers handler.Methods, log *logrus.Logger) {

	log.Info("Setting up routes...")
	log.Infof("API Documentation can be found on the path: %s", "http://127.0.0.1:8080/swagger/index.html")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	crudroutes := router.Group("/v1/swift-codes")
	crudroutes.GET("/:swift-code", handlers.GetBySwiftCode)
	crudroutes.GET("/country/:countryISO2code", handlers.GetByCountryISO2Code)

	crudroutes.POST("/", handlers.Create)
	crudroutes.DELETE("/:swift-code", handlers.DeleteBySwiftCode)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// Just for debug to make sure all methods are registered and mounted
	for _, route := range router.Routes() {
		log.Infof("Method: %s, Path: %s", route.Method, route.Path)
	}
}
