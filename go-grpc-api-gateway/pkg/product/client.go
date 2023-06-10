package product

import (
	"fmt"

	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/product/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/product/routes"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewProductServiceClient(cc)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}