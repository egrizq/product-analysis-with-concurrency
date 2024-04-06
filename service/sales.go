package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
	"strconv"
)

func ImportSalesToDatabase(csvRecords [][]string, mapProductID map[string]int) model.Response {
	for index, salesRecord := range csvRecords {
		if index != 0 {
			productId := mapProductID[salesRecord[0]]
			qty, _ := strconv.Atoi(salesRecord[1])
			date := helpers.ConvertDate(salesRecord[2])

			query := `INSERT INTO sales(product_id, qty_sold, date) VALUES (?, ?, ?)`
			if err := database.DB.Exec(query, productId, qty, date).Error; err != nil {
				return helpers.Response("Error QUERY: {INSERT INTO sales(product_id, qty_sold) VALUES (?, ?)}",
					500, err.Error())
			}
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", len(csvRecords)-1)
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}

func CountSalesProduct() {}
