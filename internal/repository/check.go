package repository

import (

	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

// get user stats 
func (r Repository) GetStats(minutes int) (int, error) {
	var count int

	query := `
		SELECT COUNT(DISTINCT user_id) 
		FROM location_checks
		WHERE created_at >= NOW() -  make_interval(mins => $1)
	`
	err := r.db.QueryRow(query, minutes).Scan(&count)

	return count, err
}

// save check request in db
func (r Repository) SaveCheck(userID string, latitude, longitude float64) (*model.LocationChecks, error) {
	var locChecks model.LocationChecks
	query := `
		INSERT INTO location_checks (user_id, lat, lng)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	err := r.db.QueryRow(
			query, userID, latitude, longitude,
		).Scan(locChecks.ID, locChecks.CreatedAt)
	
	if err != nil {
		return nil, err
	}

	return &locChecks, nil
}