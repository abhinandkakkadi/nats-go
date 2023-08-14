package main

import (
	"log"
	pb "natsgo/protobuff"
	"runtime"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

var orderServiceUri = "flipkart.com"

func main() {

	// client init
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	natsConnection.Subscribe("Discovery.OrderService", func(m *nats.Msg) {

		orderServiceDiscovery := pb.ServiceDiscovery{OrderServiceUri: orderServiceUri}
		data, err := proto.Marshal(&orderServiceDiscovery)
		if err == nil {
			natsConnection.Publish(m.Reply, data)
		}

		// keep connection alive
		runtime.Goexit()

	})
}
