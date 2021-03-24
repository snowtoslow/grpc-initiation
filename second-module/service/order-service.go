package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"grpc-initiation/second-module/service/order/orderinfo"
	"log"
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
