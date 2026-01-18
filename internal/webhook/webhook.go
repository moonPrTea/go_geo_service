package webhook

import (
	"net/http"
	"time"
)

type Queue interface {
	Pop() (string, error)
}

type Webhook struct {
	queue      Queue
	webhookURL string
	client     *http.Client
}

func NewWebhook(queue Queue, webhookURL string) *Webhook {
	return &Webhook{
		queue:      queue,
		webhookURL: webhookURL,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
