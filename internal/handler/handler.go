package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moonPrTea/go_geo_service.git/config"
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
	"github.com/moonPrTea/go_geo_service.git/internal/handler/middleware"
)

type Service interface {
	Create(ctx context.Context, req dto.IncidentRequest) (dto.IncidentResponse, error)
	GetIncidentByID(ctx context.Context, id int) (*dto.IncidentResponse, error)
	GetAllIncidents(ctx context.Context, searchActive bool) (*dto.IncidentListResponse)
	Update(ctx context.Context, id int, req dto.IncidentRequest) error
	Delete(ctx context.Context, id int) error

	// stats
	GetRequestStatistic(ctx context.Context, windowTimeMinutes int) (*dto.StatsResponse, error)

	// check location and return danger zones
	CheckLocation(ctx context.Context, req dto.CheckLocationRequest) (*dto.CheckLocationResponse, error)
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

	// auth middleware for API key validation
	auth := middleware.NewMiddleware(config.New().APIKey)


	api := r.Group("/api/v1")
	{
		incidents := api.Group("/incidents")
		incidents.Use(auth.ApiToken)
		{
			incidents.POST("/", h.CreateIncident)	
			incidents.GET("/", h.GetAllIncidents)
			incidents.GET("/:id", h.GetIncident)
			incidents.PUT("/:id", h.UpdateIncident)
			incidents.DELETE("/:id", h.DeleteIncident)

			incidents.POST("/stats", h.GetRequestStatistic)
		}

		location := api.Group("/location")
		{
			location.POST("/check", h.CheckLocation)
		}

		health := api.Group("/system")
		health.Use(auth.ApiToken)
		{
			health.GET("/health", h.HealthCheck)
		}
	}

	return r
}