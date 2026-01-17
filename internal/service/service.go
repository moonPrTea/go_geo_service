package service

import (

	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

type Repository interface {
	CreateIncident(i *model.Incident) error
}

type Service struct {
	// data repositories
	Repository
}

func New(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}