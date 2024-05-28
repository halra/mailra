package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"message": "Server is running smoothly",
		"time":    time.Now().Format(time.RFC3339),
	})
}
