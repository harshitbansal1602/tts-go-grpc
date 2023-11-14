package client

import (
	"log"

	pb "github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/bridge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetClient(serverAdd string) (pb.BridgeClient, *grpc.ClientConn) {
	// TODO: insecrue connection

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// TODO: Closing the connection should not be the responsibilty of caller
	conn, err := grpc.Dial(serverAdd, opts...)
	if err != nil {
		log.Fatalf("Failed to establish connection to: %v \n %v", serverAdd, err)
	}

	client := pb.NewBridgeClient(conn)
	return client, conn
}
