package handlers

import (
	"log"
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
		Id:       id,
		Name:     "Latte",
		Calories: 150,
	}

	log.Printf("Getting coffee with id: %s", id)

	metrics.RecordCoffeeRequest("success")
	c.JSON(http.StatusOK, resp)
}

func PutCoffee(c *gin.Context) {
	var coffee domain.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		metrics.RecordCoffeeRequest("bad_request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid coffee payload",
		})
		return
	}

	log.Printf("Creating coffee with id: %s", coffee.Id)

	metrics.RecordCoffeeRequest("success")
	c.JSON(http.StatusOK, coffee)
}
