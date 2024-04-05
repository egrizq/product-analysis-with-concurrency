package service

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"product-store-management/database"
	"product-store-management/model"
)

func ImportProductToDatabase() (model.Response, error) {
	fileJson, err := os.ReadFile("product.json")
	if err != nil {
		log.Printf("Error open JSON file: %v", err)
		return model.Response{}, err
	}

	var listProduct []model.Product

	err = json.Unmarshal(fileJson, &listProduct)
	if err != nil {
		log.Printf("Error unmarshal JSON data: %v", err)
		return model.Response{}, err
	}

	for _, product := range listProduct {
		newProduct := model.Product{
			Name:         product.Name,
			Stock:        product.Stock,
			SellingPrice: product.SellingPrice,
			BuyingPrice:  product.BuyingPrice,
		}

		err := database.DB.Create(&newProduct).Error
		if err != nil {
			log.Printf("Error insert product to database: %v", err)
			return model.Response{}, err
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data", len(listProduct))
	Response := model.Response{
		Payload:    payload,
		StatusCode: 200,
		Message:    "Import json data to database is success",
	}

	return Response, nil
}
