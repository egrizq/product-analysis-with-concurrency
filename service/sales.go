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

	mapProductID := make(map[string]int)
	for _, product := range listProduct {
		mapProductID[product.Name] = product.Id
	}

	for index, salesRecord := range csvRecords {
		if index != 0 {
			productId := mapProductID[salesRecord[0]]
			qty, _ := strconv.Atoi(salesRecord[1])

			query := `INSERT INTO sales(product_id, qty_sold) VALUES (?, ?)`
			if err := database.DB.Exec(query, productId, qty).Error; err != nil {
				return helpers.Response("Error QUERY: {INSERT INTO sales(product_id, qty_sold) VALUES (?, ?)}",
					500, err.Error())
			}
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", len(csvRecords)-1)
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}
