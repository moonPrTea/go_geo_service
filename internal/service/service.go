package service

import (
	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

type Repository interface {
	CreateIncident(i *model.Incident) error
	GetIncidentByID(id int) (*model.Incident, error)
	FindAllIncidents(searchActive bool) ([]model.Incident, error)
	FindNearbyIncidents(latitude, longitude, radius float64) ([]model.Incident, error)
	UpdateIncident(i *model.Incident) error
	DeleteIncident(id int) error

	// unique user_id for the last n minutes
	GetStats(minutes int) (int, error)

	SaveCheck(userID string, latitude, longitude float64) (*model.LocationChecks, error)
}

// redis
type Queue interface {
	Push(payload string) error
}

type Service struct {
	// data repositories
	Repository
	Queue
}

func New(repository Repository, queue Queue) Service {
	return Service{
		Repository: repository,
		Queue: queue,
	}
}