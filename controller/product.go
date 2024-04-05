package controller

import (
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessProduct(ctx *gin.Context) {
	response := service.ImportProductToDatabase()
	ctx.JSON(response.StatusCode, response)
}
