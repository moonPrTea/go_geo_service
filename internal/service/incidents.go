package service

import (
	"context"

	"github.com/moonPrTea/go_geo_service.git/internal/dto"
	"github.com/moonPrTea/go_geo_service.git/internal/mapper"
	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

func (s Service) Create(ctx context.Context, req dto.IncidentRequest) (dto.IncidentResponse, error) {
	incident := model.Incident{
        Title: req.Title,
        Latitude: req.Latitude,
        Longitude: req.Longitude,
        Radius: req.Radius,
        Active: *req.Active,
    }

    if err := s.Repository.CreateIncident(&incident); err != nil {
        return dto.IncidentResponse{}, err
    }

    return mapper.ToResponseData(incident), nil
}