package webhook

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
)

/*
service is waiting for new event, takes payload 
then send webhook
*/
func (w *Webhook) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("app is come to end, webhook work is stopped")
			return
		default:
			// waiting for new message payload
			msg, err := w.queue.Pop()
			if err != nil {
				log.Println("queue error:", err)
				continue
			}

			if err := w.send(msg); err != nil {
				log.Println("An error with webhooks occured:", err)
			}
		}
	}
}

// send webhook with json payload
func (w *Webhook) send(payload string) error {
	resp, err := w.client.Post(
		w.webhookURL,
		"application/json",
		bytes.NewBuffer([]byte(payload)),
	)
	log.Print("new webhook: ", payload)
	if err != nil {
		return err
	}

	// close connection
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("webhook work failed with status %s", resp.Status)
	}

	return nil
}