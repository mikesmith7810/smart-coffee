package router

import (
	"smart-coffee/handlers"
	"smart-coffee/metrics"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New() *gin.Engine {
	r := gin.Default()
	r.Use(metrics.Middleware())

	r.GET("/coffee/", handlers.GetCoffee)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
