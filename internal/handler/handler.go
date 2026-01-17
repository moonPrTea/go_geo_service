package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

type Service interface {
	Create(ctx context.Context, req dto.IncidentRequest) (dto.IncidentResponse, error)

}

type Handler struct {
	service Service
}

func New(service Service) Handler {
	return Handler{
		service: service,
	}
}


// initialize routing information
func (h *Handler) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	{
		incidents := api.Group("/incidents")
		{
			incidents.POST("/", h.CreateIncident)	
			incidents.GET("/",)
			incidents.GET("/:id")
			incidents.PUT("/:id", )
			incidents.DELETE("/:id", )

			incidents.POST("/stats")
		}

		location := api.Group("/location")
		{
			location.POST("/check")
		}

		health := api.Group("/system")
		{
			health.GET("/health")
		}
	}

	return r
}