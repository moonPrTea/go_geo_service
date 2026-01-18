package dto

type StatsResponse struct {
    UserCount int `json:"user_count"`
    WindowMinutes int `json:"window_minutes"`
    Timestamp string `json:"timestamp"`
}