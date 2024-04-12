package main

import (
	"product-store-management/controller"
	"product-store-management/database"
	"product-store-management/route"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connection()

	router := gin.Default()

	router.POST("/process/insert", route.InsertData)
	router.POST("/process/reports", controller.ProcessReports)

	router.Run(":8000")
}
