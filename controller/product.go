package controller

import (
	"product-store-management/helpers"
	"product-store-management/model"
	"product-store-management/service"
)

func InsertProductJSON() model.Response {
	listProduct, err := helpers.ReadJSON("public/product.json")
	if err != nil {
		return helpers.Response("Error open JSON file", 500, err.Error())
	}

	return service.ImportProductToDatabase(listProduct)
}
