package controller

import (
	"net/http"
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessProduct(ctx *gin.Context) {
	response, err := service.ImportProductToDatabase()
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, response)
}
