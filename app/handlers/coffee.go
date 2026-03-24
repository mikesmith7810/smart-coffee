package handlers

import (
	"errors"
	"net/http"
	"smart-coffee/domain"
	"smart-coffee/metrics"
	"smart-coffee/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.CoffeeService
}

func NewHandler(svc *service.CoffeeService) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) GetCoffee(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		metrics.RecordCoffeeRequest("bad_request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id parameter is required"})
		return
	}

	coffee, err := h.service.GetCoffee(id)
	if errors.Is(err, service.ErrNotFound) {
		metrics.RecordCoffeeRequest("not_found")
		c.JSON(http.StatusNotFound, gin.H{"error": "coffee not found"})
		return
	}
	if err != nil {
		metrics.RecordCoffeeRequest("error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get coffee"})
		return
	}

	metrics.RecordCoffeeRequest("success")
	c.JSON(http.StatusOK, coffee)
}

func (h *Handler) PutCoffee(c *gin.Context) {
	var coffee domain.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		metrics.RecordCoffeeRequest("bad_request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coffee payload"})
		return
	}

	if err := h.service.PutCoffee(coffee); err != nil {
		metrics.RecordCoffeeRequest("error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save coffee"})
		return
	}

	metrics.RecordCoffeeRequest("success")
	c.JSON(http.StatusOK, coffee)
}
