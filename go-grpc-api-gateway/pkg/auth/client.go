package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
)

// handler 与 grpc.Client 捆绑一下

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// WithInsecure()，没有 ssl 证书，将就用
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
