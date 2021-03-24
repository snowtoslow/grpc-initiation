package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"grpc-initiation/second-module/client/order/orderinfo"
	"io"
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


	searchStream, _ := orderMgtClient.SearchOrders(context.Background(),
		&wrappers.StringValue{Value: "google"})

	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}
		// обрабатываем другие потенциальные ошибки
		log.Print("Search Result : ", searchOrder)
	}


}
