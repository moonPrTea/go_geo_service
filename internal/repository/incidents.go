package repository

import (
	"github.com/moonPrTea/go_geo_service.git/internal/model"
)

func (r Repository) CreateIncident(i *model.Incident) error {
	err := r.db.QueryRow(`
        INSERT INTO incidents (title, lat, lng, radius, active)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at
    `, i.Title, i.Latitude, i.Longitude, i.Radius, i.Active).Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
    
    return err
}
