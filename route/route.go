package route

import (
	"product-store-management/controller"
	"product-store-management/model"

	"github.com/gin-gonic/gin"
)

func ProcessData(ctx *gin.Context) {
	var response []model.Response

	responseProduct := controller.InsertProductJSON()
	if responseProduct.StatusCode != 200 {
		ctx.JSON(responseProduct.StatusCode, responseProduct)
	}

	responseSales := controller.InsertProductCSV()
	if responseProduct.StatusCode != 200 {
		ctx.JSON(responseSales.StatusCode, responseSales)
	}

	response = append(response, responseProduct, responseSales)

	ctx.JSON(200, response)
}