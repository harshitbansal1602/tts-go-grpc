package client

import (
	"log"
	"sync"

	pb "github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/bridge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var SERVER_ADD = "localhost:50051"

type Client interface {
	Cleanup()
}

type GRPCClient struct {
	Stub pb.BridgeClient
}

var stub *GRPCClient
var conn *grpc.ClientConn
var lock = &sync.Mutex{}

func GetGRPCInstance() *GRPCClient {
	if stub == nil {
		lock.Lock()
		defer lock.Unlock()
		if stub == nil {
			opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
			// TODO: Closing the connection should not be the responsibilty of caller
			tempConn, err := grpc.Dial(SERVER_ADD, opts...)
			if err != nil {
				log.Fatalf("Failed to establish connection to: %v \n %v", SERVER_ADD, err.Error())
			}
			conn = tempConn
			stub = &GRPCClient{Stub: pb.NewBridgeClient(conn)}
			log.Println("Connection established with grpc")
		}
	}
	return stub
}

func (*GRPCClient) Cleanup() {
	if conn != nil {
		err := conn.Close()
		stub = nil
		if err != nil {
			log.Printf("Failed to close connection. %v", err.Error())
		}
		log.Println("Connection established with grpc")
	}
}
