package handler

import (
	"github.com/gin-gonic/gin"
	
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

func (h *Handler) CreateIncident(c * gin.Context) {
	var req dto.IncidentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}	

	incident, err := h.service.Create(c, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, incident)
}