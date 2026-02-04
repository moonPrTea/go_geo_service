package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/moonPrTea/go_geo_service.git/internal/dto"
)

// check location by (lat, lng) and push webhook
func (s *Service)CheckLocation (ctx context.Context, req dto.CheckLocationRequest) (*dto.CheckLocationResponse, error) {
	s.Repository.SaveCheck(req.UserId, req.Latitude, req.Longitude)

	var radius = 5.0

	incidents, err := s.Repository.FindNearbyIncidents(req.Latitude, req.Longitude, radius)
	if err != nil {
		return &dto.CheckLocationResponse{
            UserId: req.UserId,
            Latitude: req.Latitude,
            Longitude: req.Longitude,
            Zones: []string{},
            IsDanger: false,
            Timestamp: time.Now(),
        }, err
	}

	var zones []string
	for _, incident := range incidents {
		zones = append(zones, incident.Title)
	}

	// prepare data for webhook
	if len(zones) > 0 {
		payload := map[string]interface{}{
			"user_id": req.UserId, 
			"latitude": req.Latitude,
			"longitude": req.Longitude, 
			"zones": zones, 
			"timestamp": time.Now().Format(time.RFC3339),
		}

		jsonData, _ := json.Marshal(payload)
		s.Queue.Push(string(jsonData))
	}

	return &dto.CheckLocationResponse{
		UserId: req.UserId,
		Latitude: req.Latitude,
		Longitude: req.Longitude,
		Zones: zones,
		IsDanger: len(zones) > 0,
		Timestamp: time.Now(),
	}, nil

}


func (s *Service) GetRequestStatistic(ctx context.Context, windowTimeMinutes int) (*dto.StatsResponse, error) {
	userCount, err := s.Repository.GetStats(windowTimeMinutes)
	if err != nil {
		return nil, err
	}

	return &dto.StatsResponse{
		UserCount: userCount,
		WindowMinutes: windowTimeMinutes,
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil
}