package service

import (
	"encoding/csv"
	"os"
	"product-store-management/helpers"
	"product-store-management/model"
)

func ImportSalesCSVtoDatabase() model.Response {
	fileCSV, err := os.Open("/data/sales_data.csv")
	if err != nil {
		return helpers.Response("Error open csv file", 500, err.Error())
	}
	defer fileCSV.Close()

	csv.NewReader(fileCSV)
}
