package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
)


func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", svc.Register)
		auth.POST("/login", svc.Login)
	}

	return svc
}


