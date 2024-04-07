package controller

import (
	"product-store-management/helpers"
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func InsertProductJSON(ctx *gin.Context) {
	listProduct, err := helpers.ReadJSON("public/product.json")
	if err != nil {
		response := helpers.Response("Error open JSON file", 500, err.Error())
		ctx.JSON(response.StatusCode, response)
	}

	response := service.ImportProductToDatabase(listProduct)
	ctx.JSON(response.StatusCode, response)
}
