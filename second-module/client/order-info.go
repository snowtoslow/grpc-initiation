package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"grpc-initiation/second-module/client/order/orderinfo"
	"log"
)

const address = "localhost:50052"

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	orderMgtClient := orderinfo.NewOrderManagementClient(conn)

	retrievedOrder, err := orderMgtClient.GetOrder(context.Background(),
		&wrappers.StringValue{Value: "order1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

}
