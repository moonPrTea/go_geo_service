package handler

import (
	"github.com/gin-gonic/gin"
	
	"github.com/moonPrTea/go_geo_service.git/config"
)

func (h Handler) GetRequestStatistic(c *gin.Context) {
	countIDs, err := h.service.GetRequestStatistic(c, config.New().StatsWindow)
	if err != nil {
		c.IndentedJSON(500, err)
		return
	}
	c.IndentedJSON(200, countIDs)
}