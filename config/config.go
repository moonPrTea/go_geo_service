package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DbURL string
	RedisAddress string
	APIKey string
	WebhookURL  string
	StatsWindow int
}

// load env variables
func New() *Config{
	godotenv.Load("../../.env")
    
    return &Config{
        AppPort:      getEnv("APP_PORT", "8080"),
        DbURL:        getEnv("DB_URL", "dbUrl"),
        RedisAddress: getEnv("REDIS_ADDRESS", "host:port"),
        APIKey:       getEnv("API_KEY", "Key"),
        WebhookURL:   getEnv("WEBHOOK_URL", ""),
        StatsWindow:  getIntValue("STATS_TIME_WINDOW_MINUTES", 60),
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getIntValue(key string, defaultVariant int) int {
    strValue := os.Getenv(key)
    if strValue == "" {
        return defaultVariant
    }
    
    value, err := strconv.Atoi(strValue)
    if err != nil {
        return defaultVariant
    }
    return value
}