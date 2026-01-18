package dto

import "time"

// output data
type IncidentResponse struct {
	ID        int       `json:"id"`
    Title     string    `json:"title"`
    Latitude       float64   `json:"lat"`
    Longitude       float64   `json:"lng"`
    Radius    float64   `json:"radius"`
    Active    bool      `json:"active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// input data
type IncidentRequest struct {
 	Title  		string  `json:"title" binding:"required,min=1,max=100"`
    Latitude    float64 `json:"lat" binding:"required,latitude"`
    Longitude   float64 `json:"lng" binding:"required,longitude"`
    Radius 		float64 `json:"radius" binding:"required,min=1,max=1000"`
    Active 		*bool   `json:"active" binding:"required"`
}

type IncidentListResponse struct {
    Incidents []IncidentResponse `json:"incidents"`
    Total     int                `json:"total"`
}