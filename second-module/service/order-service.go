package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"grpc-initiation/second-module/service/order/orderinfo"
	"io"
	"log"
	"strings"
)

type server struct {
	orderMap map[string]*orderinfo.Order
}

func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*orderinfo.Order, error) {
	// Реализация сервиса
	if v, ok := s.orderMap[orderId.Value]; ok {
		return v, nil
	} else {
		log.Println(orderId.Value)
	}
	return nil, fmt.Errorf("error getting value from map by id: %s ", orderId.Value)
}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream orderinfo.OrderManagement_SearchOrdersServer) error {
	for key, order := range s.orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := stream.Send(order)
				if err != nil {
					return fmt.Errorf(
						"error sending message to stream : %v", err)
				}
				log.Print("Matching Order Found : " + key)
				break
			}
		}
	}
	return nil
}

func (s *server) UpdateOrders(stream orderinfo.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {

			return stream.SendAndClose(
				&wrappers.StringValue{Value: "Orders processed " + ordersStr})
		}
		s.orderMap[order.Id] = order
		log.Printf("Order ID: %s", order.Id)
		ordersStr += order.Id + ", "
		log.Printf("Updated orders: %v\n", s.orderMap)
	}
}
