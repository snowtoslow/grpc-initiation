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

	/////unary rpc example;
	retrievedOrder, err := orderMgtClient.GetOrder(context.Background(),
		&wrappers.StringValue{Value: "order1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

	////Streaming server side rpc example;
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


	/////Streaming client rpc example;
	updateStream, err := orderMgtClient.UpdateOrders(context.Background())
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", orderMgtClient, err)
	}

	if err = updateStream.Send(&orderinfo.Order{
		Id: "order1",
		Items: []string{"vova is testing update!"},
	}); err != nil {
		log.Fatalf("%v.Send() =%v", updateStream, err)
	}
	// обновляем заказ 2
	if err = updateStream.Send(&orderinfo.Order{
		Id: "order2",
		Items: []string{"myMagicRpc!"},
	}); err != nil {
		log.Fatalf("%v.Send() = %v", updateStream, err)
	}
	// обновляем заказ 3
	if err = updateStream.Send(&orderinfo.Order{
		Id: "order3",
		Items: []string{"stream 3"},
	}); err != nil {
		log.Fatalf("%v.Send() = %v", updateStream, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v",
			updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)


}
