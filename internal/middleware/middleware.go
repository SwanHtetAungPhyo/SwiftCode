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
	// HttpRequestDuration is a Prometheus histogram for tracking the duration of HTTP requests.
	// It records the duration of requests for each path, method, and response status.
	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 1, 1.5, 2, 3},
		},
		[]string{"path", "method", "status"},
	)

	// HttpRequestTotal is a Prometheus counter for tracking the total number of HTTP requests.
	// It counts the number of requests by path, method, and response status.
	HttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)
)

// init registers the Prometheus metrics to make them available for scraping.
func init() {
	prometheus.MustRegister(HttpRequestTotal)
	prometheus.MustRegister(HttpRequestDuration)
}

// PrometheusMetrics is a middleware that tracks HTTP request durations and counts for Prometheus.
// It records the request duration in seconds and the total number of requests, including status codes.
func PrometheusMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Record the start time
		path := c.FullPath()

		// If path is empty, set it to "unknown"
		if path == "" {
			path = "unknown"
		}

		// Process the request
		c.Next()

		// Calculate the duration and status of the request
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		// Record the request duration and total requests for Prometheus
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

// ResponseTimeMiddleware logs the time taken for each request.
func ResponseTimeMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start)
		log.Infof("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}

// SetUp sets up the middleware for the Gin router.
// It includes response time logging, CORS, recovery, logger, and Prometheus metrics.
func SetUp(router *gin.Engine, log *logrus.Logger) {
	router.Use(ResponseTimeMiddleware(log)) // Logs response times for each request.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Recovery())      // Recovers from panics to prevent server crashes.
	router.Use(gin.Logger())        // Logs each HTTP request.
	router.Use(PrometheusMetrics()) // Tracks request durations and counts for Prometheus.
}
