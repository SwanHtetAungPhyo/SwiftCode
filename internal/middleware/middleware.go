package middleware

import "github.com/gin-gonic/gin"

func SetUp(router *gin.Engine) {
	router.Use(func(c *gin.Context) {
		c.Header("Server", "SWIFT_CODE")
		c.Header("Author", "SwanHtetAungPhyo")
	})
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use()
}
