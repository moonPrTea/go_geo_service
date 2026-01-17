package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

type Service interface {
	Create(ctx context.Context, req dto.IncidentRequest) (dto.IncidentResponse, error)
	GetIncidentByID(ctx context.Context, id int) (*dto.IncidentResponse, error)
	GetAllIncidents(ctx context.Context, searchActive bool) (*dto.IncidentListResponse)
	Update(ctx context.Context, id int, req dto.IncidentRequest) error
	Delete(ctx context.Context, id int) error
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
			incidents.GET("/", h.GetAllIncidents)
			incidents.GET("/:id", h.GetIncident)
			incidents.PUT("/:id", h.UpdateIncident)
			incidents.DELETE("/:id", h.DeleteIncident)

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