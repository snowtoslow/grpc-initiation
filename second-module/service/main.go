package main

import (
	"google.golang.org/grpc"
	"grpc-initiation/second-module/service/order/orderinfo"
	"log"
	"net"
)

const port = ":50052"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	orderinfo.RegisterOrderManagementServer(s, &server{
		map[string]*orderinfo.Order{
			"order1": {
				Id:          "1",
				Description: "Order 1",
				Price:       30,
				Destination: "Chizdirovka noua",
				Items: []string{"google main","amazon","azure","gRPC","google"},
			},
			"order2": {
				Id:          "3",
				Description: "Order 2",
				Price:       40,
				Destination: "Chizdirovka veche",
				Items: []string{"google cloud","amazon","azure","gRPC","google-production"},
			},
			"order3": {
				Id:          "3",
				Description: "Order 3",
				Price:       30,
				Destination: "Una scurta",
				Items: []string{"google cloud","amazon","azure","gRPC"},
			},
		},
	})

	log.Printf("Starting gRPC listener on port " + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
