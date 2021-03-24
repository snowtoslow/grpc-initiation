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
			},
			"order2": {
				Id:          "3",
				Description: "Order 2",
				Price:       40,
				Destination: "Chizdirovka veche",
			},
			"order3": {
				Id:          "3",
				Description: "Order 3",
				Price:       30,
				Destination: "Una scurta",
			},
		},
	})

	log.Printf("Starting gRPC listener on port " + port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
