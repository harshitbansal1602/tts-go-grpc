package client

import (
	"log"
	"sync"

	pb "github.com/harshitbansal1602/tts-go-grpc/tts-go-grpc/bridge"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Stub pb.BridgeClient
}

var stub *GRPCClient
var conn *grpc.ClientConn
var lock = &sync.Mutex{}
var GRPC_SERVER_ADD = "localhost:50051" // TODO: Place in the env file

func GetGRPCInstance() *GRPCClient {
	if stub == nil {
		lock.Lock()
		defer lock.Unlock()
		if stub == nil {
			stub = stub.init()
		}
	}
	return stub
}

func (*GRPCClient) init() *GRPCClient {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	tempConn, err := grpc.Dial(GRPC_SERVER_ADD, opts...)
	if err != nil {
		log.Fatalf("Failed to establish connection to: %v \n %v", GRPC_SERVER_ADD, err.Error())
	}
	conn = tempConn
	stub := &GRPCClient{Stub: pb.NewBridgeClient(conn)}
	log.Println("Connection established with grpc")
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
