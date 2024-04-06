package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
)

func ImportProductToDatabase() model.Response {
	listProduct, err := helpers.ReadJSON("public/product.json")
	if err != nil {
		return helpers.Response("Error open JSON file", 500, err.Error())
	}

	// create a hashmap to check if the product exists
	seenProducts := make(map[string]bool)

	for _, product := range listProduct {
		if !seenProducts[product.Name] {
			query := `INSERT INTO products(name, stock, selling_price, buying_price) VALUES (?, ?, ?, ?)`

			err := database.DB.Exec(query, product.Name, product.Stock, product.SellingPrice, product.BuyingPrice).Error
			if err != nil {
				return helpers.Response("Error insert product to database", 500, err.Error())
			}

			seenProducts[product.Name] = true
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", len(seenProducts))
	response := helpers.Response(payload, 200, "Import json data to database is success")

	return response
}
