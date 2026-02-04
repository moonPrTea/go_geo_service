package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// read request body text
		body, _ := io.ReadAll(r.Body)
		log.Println("webhook:", string(body))

		// send feedback status 200
		w.WriteHeader(http.StatusOK)
	})

	log.Println("webhook listening on :9090 port")
	
	// waiting for all request
	log.Fatal(http.ListenAndServe(":9090", nil))
}
