package service

import (
	"github.com/moonPrTea/go_geo_service.git/internal/repository"
)

// redis
type Queue interface {
	Push(payload string) error
}

type Service struct {
	// data repositories
	*repository.Repository
	Queue
}

func New(repository *repository.Repository, queue Queue) *Service {
	return &Service{
		Repository: repository,
		Queue: queue,
	}
}