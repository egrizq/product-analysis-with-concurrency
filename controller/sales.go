package controller

import (
	"product-store-management/helpers"
	"product-store-management/service"

	"github.com/gin-gonic/gin"
)

func ProcessSales(ctx *gin.Context) {
	csvRecords, err := helpers.ReadCSV("public/sales_data.csv")
	if err != nil {
		response := helpers.Response("Error read csv file", 500, err.Error())
		ctx.JSON(response.StatusCode, response)
	}

	response := service.ImportSalesToDatabase(csvRecords)
	ctx.JSON(response.StatusCode, response)
}
