package handlers

import (
	"net/http"
	"smart-coffee/domain"
	"smart-coffee/metrics"

	"github.com/gin-gonic/gin"
)

func GetCoffee(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		metrics.RecordCoffeeRequest("bad_request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id parameter is required",
		})
		return
	}

	resp := domain.Coffee{
		Id:       "1",
		Name:     "Latte",
		Calories: 150,
	}

	metrics.RecordCoffeeRequest("success")
	c.JSON(http.StatusOK, resp)
}
