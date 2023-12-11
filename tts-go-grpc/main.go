package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/client"
	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/handlers"
)

func main() {
	client := client.GetGRPCInstance()
	defer client.Cleanup()

	http.HandleFunc("/init", handlers.InitModels)
	http.HandleFunc("/speech", handlers.GetSpeech)
	log.Printf("started server")

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
