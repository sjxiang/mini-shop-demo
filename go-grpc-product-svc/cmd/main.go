package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/mini-shop-demo/go-grpc-product-svc/pkg/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-product-svc/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-product-svc/pkg/db"
	"github.com/sjxiang/mini-shop-demo/go-grpc-product-svc/pkg/service"
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

	fmt.Println("Product svc on ", c.Port)

	s := service.Server{
		H:   h,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("启动失败：", err)
	}

}