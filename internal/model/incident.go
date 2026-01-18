package model

import "time"

type Incident struct {
	ID        int
	Title     string
	Latitude  float64
	Longitude float64
	Radius    float64
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}


