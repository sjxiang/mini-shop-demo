package service

import (
	"context"
	"net/http"

	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/db"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/model"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/pb"
	"github.com/sjxiang/mini-shop-demo/go-grpc-auth-svc/pkg/util"
)

type Server struct {
	Jwt util.JwtWrapper
	H   db.Handler

	pb.UnimplementedAuthServiceServer
}


func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user model.User

	if result := s.H.DB.Where(&model.User{Email: req.Email}).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error: "Email already exists.",
		}, nil
	}

	user.Email = req.Email
	user.Password = util.HashPassowrd(req.Password)

	s.H.DB.Create(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}


func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user model.User

	if result := s.H.DB.Where(&model.User{Email: req.Email}).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound, 
			Error: "user not found.",
		}, nil
	}

	match := util.CheckPassowrdHash(user.Password, req.Password)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound, 
			Error: "user not found.",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token: token,
	}, nil
}


func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error: err.Error(),
		}, nil 
	}	

	var user model.User

	if result := s.H.DB.Where(&model.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound, 
			Error: "user not found.",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}