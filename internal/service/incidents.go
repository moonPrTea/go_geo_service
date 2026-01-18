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

func (s Service) Update(ctx context.Context, id int, req dto.IncidentRequest) error {
    incident, err := s.Repository.GetIncidentByID(id)
    if err != nil {
        return err
    }


    incident.Title = req.Title
    incident.Latitude = req.Latitude
    incident.Longitude = req.Longitude
    incident.Radius = req.Radius
    incident.Active = *req.Active

    return s.Repository.UpdateIncident(incident)
}

func (s Service) Delete(ctx context.Context, id int) error {
    return s.Repository.DeleteIncident(id)
}

func (s Service) GetIncidentByID(ctx context.Context, id int) (*dto.IncidentResponse, error) {
	incident, err := s.Repository.GetIncidentByID(id)
    if err != nil {
        return nil, err
    }

    response := mapper.ToResponseData(*incident)
    return &response, nil
}

func (s Service) GetAllIncidents(ctx context.Context, searchActive bool) *dto.IncidentListResponse {
    incidents, err := s.Repository.FindAllIncidents(searchActive)
    if err != nil {
        return &dto.IncidentListResponse{}
    }

    var response dto.IncidentListResponse

    // convert values and append them to result value
    for _, incident := range incidents {
        response.Incidents = append(response.Incidents, mapper.ToResponseData(incident))
    }

    // set total num
    response.Total = len(response.Incidents)

    return &response
}