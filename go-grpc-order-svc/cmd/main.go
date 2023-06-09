package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/client"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/db"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/service"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("加载配置失败:", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("监听失败：", err)
	}
	fmt.Println("订单服务开启：", c.Port)

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	s := service.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("启动失败：", err)
	}

}