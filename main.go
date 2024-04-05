package main

import (
	"product-store-management/controller"
	"product-store-management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectionDB()

	router := gin.Default()

	router.POST("/process/product", controller.ProcessProduct)

	router.Run(":8000")
}
