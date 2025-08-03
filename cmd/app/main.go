package main

import (
	"github.com/gin-gonic/gin"
	"github.com/latestcomment/go-gin-web-server/internal/db"
	"github.com/latestcomment/go-gin-web-server/internal/handlers"
)

func main() {
	db.InitDB()
	router := gin.Default()

	router.GET("/customers", handlers.GetAllCustomers)
	router.GET("/customer/:id", handlers.GetCustomerByID)
	router.PUT("/customer/update", handlers.UpdateCustomerByID)

	router.Run("localhost:8081")
}
