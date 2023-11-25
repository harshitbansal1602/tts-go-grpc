package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/bridge"
	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/client"
	"google.golang.org/protobuf/types/known/emptypb"
)

func InitModels(w http.ResponseWriter, r *http.Request) {
	client := client.GetGRPCInstance()

	_, err := client.Stub.DownloadBarkModel(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf(err.Error())
		http.Error(w, "Failed to start models", http.StatusFailedDependency)
	}
	log.Printf("models initiated")
	io.WriteString(w, "Models initiated.")
}

func GetSpeech(w http.ResponseWriter, r *http.Request) {
	client := client.GetGRPCInstance()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
	}

	var text bridge.Text
	err = json.Unmarshal(body, &text)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	speech, err := client.Stub.GetSpeech(context.Background(), &text)
	if err != nil {
		log.Fatalf("Failed to get speech. \n %v", err)
		http.Error(w, "failed to get speech", http.StatusInternalServerError)
		return
	}
	outputFile, err := os.Create("speech.wav")
	if err != nil {
		log.Fatalf("Failed to create file speech.wav")
		return
	} else {
		_, err = outputFile.Write(speech.Speech)
		if err != nil {
			log.Fatalf("failed to write to file")
		}
	}
	log.Printf("made file")
	w.Write(speech.Speech)
}
