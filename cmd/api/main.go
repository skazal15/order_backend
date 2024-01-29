package main

import (
	"order-service/internal/app/database"
	"order-service/internal/app/order/handler"
	"order-service/internal/app/order/repository"
	"order-service/internal/app/order/service"
	"order-service/public/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	route := gin.Default()

	db := database.NewDatabaseConn()
	orderRepo := repository.NeworderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	configRoute := cors.DefaultConfig()
	configRoute.AllowAllOrigins = true
	configRoute.AllowMethods = config.CORSAllowedMethods
	configRoute.AllowHeaders = config.CORSAllowedHeader

	groupRoute := route.Group("/order").Use(cors.New(configRoute))
	{
		groupRoute.GET("/", orderHandler.GetOrder)
		groupRoute.GET("/search", orderHandler.GetOrderByOrderORProduct)
		groupRoute.GET("/date", orderHandler.GetOrderByDateRange)
	}

	route.Run(":" + config.PORTApp)
}
