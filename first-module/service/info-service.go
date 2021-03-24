package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc-initiation/first-module/service/ecommerce/productinfo"
)

type server struct {
	productMap map[string]*productinfo.Product
}

// AddProduct реализует ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, in *productinfo.Product) (*productinfo.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*productinfo.Product)
	}
	s.productMap[in.Id] = in
	return &productinfo.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

// GetProduct реализует ecommerce.GetProduct
func (s *server) GetProduct(ctx context.Context, in *productinfo.ProductID) (*productinfo.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}
