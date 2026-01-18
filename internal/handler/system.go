package handler

import (
	"time"

	"github.com/gin-gonic/gin"
)


func (h Handler) HealthCheck(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
			"status": "ok",
			"system": "active", 
			"checked_at": time.Now().Unix(),
	})
}