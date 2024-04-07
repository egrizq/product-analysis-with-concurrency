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

	listProduct, err := service.GetProductNameAndID()
	if err != nil {
		response := helpers.Response("Error QUERY: {SELECT id, name FROM products;}", 500, err.Error())
		ctx.JSON(response.StatusCode, response)
	}

	// hashmap to build table relation with product.id from sales table
	mapProductID := make(map[string]int)
	for _, product := range listProduct {
		mapProductID[product.Name] = product.Id
	}

	response := service.ImportSalesToDatabase(csvRecords, mapProductID)
	ctx.JSON(response.StatusCode, response)
}
