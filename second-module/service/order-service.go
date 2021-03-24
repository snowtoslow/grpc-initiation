package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"grpc-initiation/second-module/service/order/orderinfo"
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
