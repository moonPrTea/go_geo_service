package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

func (h *Handler) CreateIncident(c * gin.Context) {
	var req dto.IncidentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}	

	incident, err := h.service.Create(c, req)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(201, incident)
}

func (h *Handler) UpdateIncident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid value for id"})
		return
	}

	var req dto.IncidentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	

	if err := h.service.Update(c, id, req); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Incident data have successfully updated"})
}

func (h *Handler) DeleteIncident(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid value for id"})
		return
	}

	if err := h.service.Delete(c, id); err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "Incident have successfully deleted"})

}

func (h *Handler) GetIncident(c * gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid value for id"})
		return
	}	

	response, err := h.service.GetIncidentByID(c, id)
	if err != nil {
		c.IndentedJSON(404, gin.H{"error": "No incident was found"})
		return
	}
	c.IndentedJSON(200, response)
}

func (h *Handler) GetAllIncidents(c *gin.Context) {
	searchActive := c.Query("active") == "true"
	response := h.service.GetAllIncidents(c, searchActive)
	c.IndentedJSON(200, response)
}