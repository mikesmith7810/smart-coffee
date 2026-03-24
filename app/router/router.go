package router

import (
	"smart-coffee/handlers"
	"smart-coffee/metrics"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func New(h *handlers.Handler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(metrics.Middleware())

	r.GET("/coffee/", h.GetCoffee)
	r.PUT("/coffee/", h.PutCoffee)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
