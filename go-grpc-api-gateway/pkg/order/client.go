package order

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/order/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/order/routes"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}


func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}