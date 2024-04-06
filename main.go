package main

import (
	"product-store-management/controller"
	"product-store-management/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
	database.Connection()

	router := gin.Default()

	router.POST("/process/product", controller.InsertProductJSON)
	router.POST("/process/sales", controller.ProcessSales)

	router.Run(":8000")
}
