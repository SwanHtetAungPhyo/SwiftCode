package routes

import (
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetUpRoute(router *gin.Engine, handlers handler.Methods, log *logrus.Logger) {
	log.Info("Setting up routes...")
	log.Infof("API Documentation can be found on the path: %s", "http://127.0.0.1:8080/swagger/index.html")

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	crudroutes := router.Group("/v1/swift-codes")
	crudroutes.GET("/:swift-code", handlers.GetBySwiftCode)
	crudroutes.POST("/", handlers.Create)
	crudroutes.DELETE("/:swift-code", handlers.DeleteBySwiftCode)
	crudroutes.GET("/country/:countryISO2code", handlers.GetByCountryISO2Code)

	for _, route := range router.Routes() {
		log.Infof("Method: %s, Path: %s", route.Method, route.Path)
	}
}
