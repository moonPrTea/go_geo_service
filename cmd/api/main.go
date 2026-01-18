package main

import (
	"context"

	"github.com/moonPrTea/go_geo_service.git/config"
	"github.com/moonPrTea/go_geo_service.git/internal/handler"
	"github.com/moonPrTea/go_geo_service.git/internal/repository"
	"github.com/moonPrTea/go_geo_service.git/internal/service"
	"github.com/moonPrTea/go_geo_service.git/internal/webhook"
	"github.com/moonPrTea/go_geo_service.git/pkg/postgres"
	"github.com/moonPrTea/go_geo_service.git/pkg/redis"
)


func main() {
	conf := config.New()

	db := postgres.Init(conf.DbURL)
	q := redis.Init(conf.RedisAddress)

	repository := repository.New(db)
	service := service.New(repository, q)
	handler := handler.New(service)

	webhook := webhook.NewWebhook(q, conf.WebhookURL)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go webhook.Run(ctx)


	r := handler.InitRouter()
	r.Run()
}