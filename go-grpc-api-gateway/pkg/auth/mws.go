package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth/pb"
)

// 认证中间件
type AuthMiddlewareBuilder struct {
	svc *ServiceClient
}

func InitAuthMiddlewareBuilder(svc *ServiceClient) AuthMiddlewareBuilder {
	return AuthMiddlewareBuilder{svc: svc}
}


func (m AuthMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 客户端携带 Token 有三种方式，具体取决于实际业务情况
		// 1. 放在请求 header
		// 2. 放在请求 body 
		// 3. 放在 url 中
		//
		// 方式一，具体做法：
		// token 放在 Header 的 Authorization 中，例如 "bearer xxx.xxx.xxx"
		
		authHeader := ctx.Request.Header.Get("Authorization")
		
		if authHeader == "" {
			// "请求 header 中 auth 为空"
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 按空格分割
		token := strings.Split(authHeader, "bearer ")
		
		if len(token) < 2 {
			// 请求 header 中，auth 格式有误
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		
		res, err := m.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
			Token: token[1],
		})

		if err != nil || res.Status != http.StatusOK {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user_id", res.UserId)

		ctx.Next()
	}
}


