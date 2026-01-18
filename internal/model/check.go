package model

import (
	"time"
)


type LocationChecks struct {
    ID        int       `db:"id"`
    UserId     string    `db:"user_id"`
    Latitude  float64   `db:"lat"`
    Longitude float64   `db:"lng"`
    CreatedAt time.Time `db:"created_at"`
}