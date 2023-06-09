package main

import (
	"log"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/auth"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/order"
	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/product"
)

func main() {
	
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("加载配置错误：", err.Error())
	}

	r := gin.Default()

	authSvc := auth.RegisterRoutes(r, &conf)
	product.RegisterRoutes(r, &conf, authSvc)
	order.RegisterRoutes(r, &conf, authSvc)

	r.Run(conf.Port)
	
}



