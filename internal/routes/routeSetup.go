package routes

import (
	_ "github.com/SwanHtetAungPhyo/swifcode/docs"
	"github.com/SwanHtetAungPhyo/swifcode/internal/handler"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func SetUpRoute(router *gin.Engine, handlers handler.Methods, log *logrus.Logger) {

	log.Info("Setting up routes...")
	log.Infof("API Documentation can be found on the path: %s", "http://127.0.0.1:8080/swagger/index.html")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cacheStore := persistence.NewInMemoryStore(time.Minute)

	crudroutes := router.Group("/v1/swift-codes")
	// These Routes are cached cuz they are the data that cannot be modified frequently so that I can reduce the IO operation on Database
	crudroutes.GET("/:swift-code", cache.CachePage(cacheStore, time.Minute*5, handlers.GetBySwiftCode))
	crudroutes.GET("/country/:countryISO2code", cache.CachePage(cacheStore, time.Minute*5, handlers.GetByCountryISO2Code))

	crudroutes.POST("/", handlers.Create)
	crudroutes.DELETE("/:swift-code", handlers.DeleteBySwiftCode)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	for _, route := range router.Routes() {
		log.Infof("Method: %s, Path: %s", route.Method, route.Path)
	}
}
