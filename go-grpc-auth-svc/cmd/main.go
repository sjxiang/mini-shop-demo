package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/config"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/db"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/service"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/util"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("加载配置失败:", err)
	}

	h := db.Init(c.DBUrl)

	jwt := util.JwtWrapper{
		SecretKey:      c.JWTSecretKey,
		Issuer:         "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("监听失败：", err)
	}

	fmt.Println("Auth svc on ", c.Port)

	s := service.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("启动失败：", err)
	}

}