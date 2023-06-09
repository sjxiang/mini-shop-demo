package service

import (
	"context"
	"net/http"

	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/client"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/db"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/model"
	"github.com/sjxiang/mini-shop-demo/go-grpc-order-svc/pkg/pb"
)

type Server struct {
	H          db.Handler
	ProductSvc client.ProductServiceClient

	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	product, err := s.ProductSvc.FindOne(req.ProductId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil 
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: product.Status,
			Error:  err.Error(),
		}, nil 
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "库存不足",
		}, nil 
	}

	order := model.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,		
	}
	
	s.H.DB.Create(&order)
	
	res, err := s.ProductSvc.DecreaseStock(req.ProductId, order.Id)
	
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(new(model.Order), order.Id)
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil 
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}