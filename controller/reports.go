package controller

import (
	"product-store-management/database"
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessReports(ctx *gin.Context) {
	start := service.Reports(database.DB)
	reports := start.ReportsRequirement()
	response := start.InsertReports(reports)

	ctx.JSON(200, response)
}
