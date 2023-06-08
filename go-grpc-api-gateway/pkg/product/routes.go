package product

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	product := r.Group("/product")
	product.Use(auth.InitAuthMiddlewareBuilder(authSvc).Build())

	{
		product.POST("/add", svc.CreateProduct)
		product.GET("/:id", svc.FindOne)
	}

}
