package main

import (
	"log"
	pb "natsgo/protobuff"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func main() {

	// client init
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to" + nats.DefaultURL)

	// send a request to subject Discovery.OrderService
	msg, err := natsConnection.Request("Discovery.OrderService", nil, 1000*time.Millisecond)
	if err == nil && msg != nil {
		orderServiceDiscovery := pb.ServiceDiscovery{}
		err := proto.Unmarshal(msg.Data, &orderServiceDiscovery)
		if err != nil {
			log.Fatalf("Error on un marshall: %v", err)
		}

		address := orderServiceDiscovery.OrderServiceUri
		log.Println("OrderService endpoint found at:", address)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		_, _ = conn, err

	}
}
