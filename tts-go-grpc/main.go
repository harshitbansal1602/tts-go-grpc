package main

import (
	"context"
	"log"
	"os"

	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/bridge"
	"github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/client"
	"google.golang.org/protobuf/types/known/emptypb"
)

var SERVER_ADD = "localhost:50051"

func main() {
	client, conn := client.GetClient(SERVER_ADD)
	defer conn.Close()

	client.DownloadBarkModel(context.Background(), &emptypb.Empty{})
	speech, err := client.GetSpeech(context.Background(), &bridge.Text{
		Text: `Hello world! This is a test sentence.
				Repeat after beep [beep] [beep] [PAUSE] [beep]
				Haha Got You!`,
	})
	if err != nil {
		log.Fatalf("Failed to get speech. \n %v", err)
		return
	}
	outputFile, err := os.Create("speech.wav")
	if err != nil {
		log.Fatalf("Failed to create file speech.wav")
		return
	}
	_, err = outputFile.Write(speech.Speech)
	if err != nil {
		log.Fatalf("Failed to write to file speech.wav")
		return
	}
	log.Printf("Called the functions")
}
