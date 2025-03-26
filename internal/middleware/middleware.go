package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func ResponseTimeMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Infof("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}
func SetUp(router *gin.Engine, log *logrus.Logger) {
	router.Use(func(c *gin.Context) {
		c.Header("Server", "SWIFT_CODE")
		c.Header("Author", "SwanHtetAungPhyo")
	})
	router.Use(ResponseTimeMiddleware(log))
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use()
}
