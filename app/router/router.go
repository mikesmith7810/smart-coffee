package router

import (
	"smart-coffee/handlers"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.GET("/coffee/", handlers.GetCoffee)

	return r
}
