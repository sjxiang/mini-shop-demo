package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/pb"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ProductServiceClient{
		Client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(productId, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}