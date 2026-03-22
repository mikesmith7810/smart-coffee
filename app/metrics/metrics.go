package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "smart_coffee_http_requests_total",
			Help: "Total number of HTTP requests handled by the service.",
		},
		[]string{"method", "route", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "smart_coffee_http_request_duration_seconds",
			Help:    "Duration of HTTP requests handled by the service.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	coffeeRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "smart_coffee_coffee_requests_total",
			Help: "Total number of coffee endpoint requests by result.",
		},
		[]string{"result"},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration, coffeeRequestsTotal)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}

		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		duration := time.Since(start).Seconds()

		httpRequestsTotal.WithLabelValues(method, route, status).Inc()
		httpRequestDuration.WithLabelValues(method, route).Observe(duration)
	}
}

func RecordCoffeeRequest(result string) {
	coffeeRequestsTotal.WithLabelValues(result).Inc()
}
