package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

func (h Handler) CheckLocation(c *gin.Context) {
	var req dto.CheckLocationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response , _ := h.service.CheckLocation(c, req)
	c.JSON(200, response)
} 