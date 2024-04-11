package controller

import (
	"product-store-management/helpers"
	"product-store-management/model"
	"product-store-management/service"
)

func InsertSalesCSV() model.Response {
	csvRecords, err := helpers.ReadCSV("public/sales_data.csv")
	if err != nil {
		return helpers.Response("Error read csv file", 500, err.Error())
	}

	return service.ImportSalesToDatabase(csvRecords)
}
