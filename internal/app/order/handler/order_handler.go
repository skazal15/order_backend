package handler

import (
	"order-service/internal/app/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService interfaces.OrderService
}

func NewOrderHandler(orderService interfaces.OrderService) *OrderHandler {
	return &OrderHandler{orderService}
}

func (oh OrderHandler) GetOrder(ctx *gin.Context) {
	pageNumber, _ := strconv.Atoi(ctx.Query("pageNumber"))
	resps, errs := oh.orderService.GetOrder(pageNumber)
	if errs != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid get order",
		})
		return
	}
	ctx.JSON(200, resps)
}

func (oh OrderHandler) GetOrderByOrderORProduct(ctx *gin.Context) {
	order := ctx.DefaultQuery("order", "")
	product := ctx.DefaultQuery("product", "")
	pageNumber, _ := strconv.Atoi(ctx.Query("pageNumber"))
	resp, err := oh.orderService.GetOrderByOrderAndProduct(pageNumber, order, product)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "fail to get data by order or product",
		})
		return
	}
	ctx.JSON(200, resp)
}

func (oh OrderHandler) GetOrderByDateRange(ctx *gin.Context) {
	start := ctx.DefaultQuery("start", "")
	end := ctx.DefaultQuery("end", "")
	pageNumber, _ := strconv.Atoi(ctx.Query("pageNumber"))
	resp, err := oh.orderService.GetOrderByDateRange(pageNumber, start, end)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "fail to get data by date",
		})
		return
	}
	ctx.JSON(200, resp)
}
