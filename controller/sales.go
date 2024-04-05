package controller

import (
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func InsertCSVFile(ctx *gin.Context) {
	response := service.ImportSalesToDatabase()
	ctx.JSON(response.StatusCode, response)
}
