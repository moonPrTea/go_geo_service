package dto

import "time"

// output data
type CheckLocationResponse struct {
	UserId     string       `json:"user_id"`
    Title      string    `json:"title"`
    Latitude   float64   `json:"lat"`
    Longitude  float64   `json:"lng"`
	Zones 	   []string  `json:"zones"`
    IsDanger   bool      `json:"is_danger"`
    Timestamp  time.Time `json:"timestamp"`
}

// input data
type CheckLocationRequest struct {
    Latitude    float64 `json:"lat" binding:"required,latitude"`
    Longitude   float64 `json:"lng" binding:"required,longitude"`
    UserId 		string  `json:"user_id"`
}
