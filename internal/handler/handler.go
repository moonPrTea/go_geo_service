package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moonPrTea/go_geo_service.git/config"
	"github.com/moonPrTea/go_geo_service.git/internal/handler/middleware"
	"github.com/moonPrTea/go_geo_service.git/internal/service"
)


type Handler struct {
	service *service.Service
}


func New(s *service.Service) *Handler {
	return &Handler{service: s}
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