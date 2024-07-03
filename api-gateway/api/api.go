package api

import (
	"api-gateway/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn,connP,orderConn *grpc.ClientConn) *gin.Engine {
	
	router := gin.Default()

	userHandler := handler.NewUserHandler(conn)
	productHandler := handler.NewProductHandler(connP)
	orderHandler := handler.NewOrderHandler(orderConn) 


	// User routes
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUser)
	router.PUT("/users", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)


	// Product routes
	router.POST("/products", productHandler.AddProduct)
	router.GET("/products/:id", productHandler.GetProduct)
	router.PUT("/products", productHandler.UpdateProduct)
	router.DELETE("/delete/:id", productHandler.DeleteProduct)
	router.GET("/products", productHandler.ListProducts)


	// Order routes
	router.POST("/orders", orderHandler.CreateOrder)
	router.GET("/orders/:id", orderHandler.GetOrder)
	router.DELETE("/orders/:id", orderHandler.DeleteOrder)
	router.GET("/orders", orderHandler.ListOrders)
	router.POST("/orders/bulk", orderHandler.CreateOrders)

	return router
}
