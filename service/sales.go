package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
	"strconv"
)

func ImportSalesToDatabase() model.Response {
	csvRecords, err := helpers.ReadCSV("public/sales_data.csv")
	if err != nil {
		return helpers.Response("Error read csv file", 500, err.Error())
	}

	var listProduct []model.ProductNameId
	query := "SELECT id, name FROM products;"
	if err := database.DB.Raw(query).Scan(&listProduct).Error; err != nil {
		return helpers.Response("Error QUERY: {SELECT id, name FROM products;}", 500, err.Error())
	}

	for index, salesRecord := range csvRecords {
		if index != 0 {
			for _, product := range listProduct {
				if salesRecord[0] == product.Name {
					qty, _ := strconv.Atoi(salesRecord[1])

					salesModel := &model.Sales{
						ProductId: product.Id,
						QtySold:   qty,
					}

					if err := database.DB.Create(salesModel).Error; err != nil {
						return helpers.Response("Error ORM: {database.DB.Create(model)}", 500, err.Error())
					}
				}
			}
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", len(csvRecords))
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}
