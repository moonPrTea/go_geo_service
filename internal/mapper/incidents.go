package mapper

import (
	"github.com/moonPrTea/go_geo_service.git/internal/dto"
	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

func ToResponseData(incident model.Incident) dto.IncidentResponse {
    return dto.IncidentResponse{
        ID:        incident.ID,
        Title:     incident.Title,
        Latitude:       incident.Latitude,
        Longitude:       incident.Longitude,
        Radius:    incident.Radius,
        Active:    incident.Active,
        CreatedAt: incident.CreatedAt,
        UpdatedAt: incident.UpdatedAt,
    }
}