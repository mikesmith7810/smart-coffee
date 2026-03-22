package handlers

import (
	"net/http"
	"smart-coffee/domain"

	"github.com/gin-gonic/gin"
)

func GetCoffee(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "id parameter is required",
        })
        return
    }

	resp := domain.Coffee{
        Id: "1",
		Name: "Latte",
		Calories: 150,
    }
    c.JSON(http.StatusOK, resp)
}
