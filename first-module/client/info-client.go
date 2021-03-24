package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-initiation/first-module/client/ecommerce/productinfo"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := productinfo.NewProductInfoClient(conn)
	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.AddProduct(ctx,
		&productinfo.Product{Name: name, Description: description})

	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)
	product, err := c.GetProduct(ctx, &productinfo.ProductID{Value: r.Value})

	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s ", product.String())
}
