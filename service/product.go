package service

import (
	"encoding/json"
	"fmt"
	"os"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
)

func ImportProductToDatabase() model.Response {
	fileJson, err := os.ReadFile("/data/product.json")
	if err != nil {
		return helpers.Response("Error open JSON file", 500, err.Error())
	}

	var listProduct []model.Product

	err = json.Unmarshal(fileJson, &listProduct)
	if err != nil {
		return helpers.Response("Error unmarshal JSON data", 500, err.Error())
	}

	for _, product := range listProduct {
		query := `INSERT INTO products(name, stock, selling_price, buying_price) VALUES (?, ?, ?, ?)`

		err := database.DB.Exec(query, product.Name, product.Stock, product.SellingPrice, product.BuyingPrice).Error
		if err != nil {
			return helpers.Response("Error insert product to database", 500, err.Error())
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data", len(listProduct))
	response := helpers.Response(payload, 200, "Import json data to database is success")

	return response
}
