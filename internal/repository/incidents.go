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

func (r Repository) UpdateIncident(i *model.Incident) error {
	err := r.db.QueryRow(`
		UPDATE incidents
		SET title = $1, lat = $2, lng = $3, radius = $4, active = $5, updated_at = NOW()
		WHERE id = $6
	`, i.Title, i.Latitude, i.Longitude, i.Radius, i.Active, i.ID,
	).Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
	
	return err
}

func (r Repository) DeleteIncident(id int) error {
    _, err := r.db.Exec(`
		DELETE FROM incidents WHERE id = $1
	`, id)
	return err
}

func (r Repository) GetIncidentByID(id int) (*model.Incident, error) {
    var incident model.Incident
	query := `
        SELECT id, title, lat, lng, radius, active, created_at, updated_at
        FROM incidents WHERE id = $1
    `
	err := r.db.QueryRow(
		query, id,
	).Scan(
		&incident.ID, &incident.Radius, &incident.Latitude, 
		&incident.Longitude, &incident.Radius, &incident.Active,
		&incident.CreatedAt, &incident.UpdatedAt, 
	)

	if err != nil {
		return nil, err
	}
	return &incident, nil
}

func (r Repository) FindAllIncidents(searchActive bool) ([]model.Incident, error) {
	query := `
	SELECT id, title, lat, lng, radius, active, created_at, updated_at
	FROM incidents
	`
	
	if searchActive {
		query += `WHERE active = true`
	}

	incidentRows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	// close cursor
	defer incidentRows.Close()

	var incidents []model.Incident
	for incidentRows.Next() {
		var i model.Incident

		// Scan reads data from 1 row and save it in current value
		err := incidentRows.Scan(
			&i.ID, &i.Title, &i.Latitude, &i.Longitude,
            &i.Radius, &i.Active, &i.CreatedAt, &i.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		incidents = append(incidents, i)
	}
	return incidents, nil
}