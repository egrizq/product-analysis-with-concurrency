package controller

import (
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessReports(ctx *gin.Context) {
	// reports := service.Reports(database.DB)

	sales := service.SalesReports()

	ctx.JSON(200, sales)
}
