package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var (
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 1.5, 2, 3},
		},
		[]string{"path", "method", "status"},
	)

	HttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
)

func init() {
	prometheus.MustRegister(HttpRequestTotal)
	prometheus.MustRegister(HttpRequestDuration)
}

// PrometheusMetrics Middleware to measure the latency and counts
func PrometheusMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()

		if path == "" {
			path = "unknown"
		}

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		HttpRequestDuration.WithLabelValues(
			path,
			c.Request.Method,
			strconv.Itoa(status),
		).Observe(duration)

		HttpRequestTotal.WithLabelValues(
			path,
			c.Request.Method,
			strconv.Itoa(status),
		).Inc()
	}
}
func ResponseTimeMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Infof("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}

func SetUp(router *gin.Engine, log *logrus.Logger) {
	router.Use(ResponseTimeMiddleware(log))
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(PrometheusMetrics())
}
