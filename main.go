package main

import (
	"product-store-management/controller"
	"product-store-management/database"
	"product-store-management/route"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
	database.Connection()

	router := gin.Default()

	router.POST("/process/insert", route.InsertData)
	router.GET("/process/reports", controller.ProcessReports)

	router.Run(":8000")
}
