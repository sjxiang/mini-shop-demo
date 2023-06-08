package routes

import (
	"net/http"
	"context"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/mini-shop-demo/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("user_id")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}