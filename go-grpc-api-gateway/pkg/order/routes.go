package order

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
)


func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	order := r.Group("/order")
	order.Use(auth.InitAuthMiddlewareBuilder(authSvc).Build())

	{
		order.POST("/add", svc.CreateOrder)
	}
	
}


