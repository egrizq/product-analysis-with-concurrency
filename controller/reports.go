package controller

import (
	"product-store-management/database"
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessReports(ctx *gin.Context) {
	process := service.ReportsInit(database.DB)
	response := process.Insert()

	ctx.JSON(response.StatusCode, response)
}
