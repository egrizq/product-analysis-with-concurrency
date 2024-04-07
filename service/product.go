package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
	"sync"
)

func removeDuplicate(batchSize int, listProduct []model.Product, wg *sync.WaitGroup) []model.Product {
	var newListProduct []model.Product

	seenProducts := make(map[string]bool)
	rangeListProduct := len(listProduct)

	for i := 0; i < rangeListProduct; i += batchSize {
		end := i + batchSize
		if end > rangeListProduct {
			end = rangeListProduct
		}

		wg.Add(1)
		batch := listProduct[i:end]

		go func(product []model.Product) {
			defer wg.Done()

			for _, product := range listProduct {
				if !seenProducts[product.Name] {
					newListProduct = append(newListProduct, product)
					seenProducts[product.Name] = true
				}
			}
		}(batch)
	}
	wg.Wait()

	return newListProduct
}

func ImportProductToDatabase(listProduct []model.Product) model.Response {
	var wg sync.WaitGroup

	batchSize := 500
	errorChannel := make(chan error)

	newListProduct := removeDuplicate(batchSize, listProduct, &wg)

	wg.Add(1)

	go func(product []model.Product) {
		defer wg.Done()

		if err := database.DB.Create(&product).Error; err != nil {
			errorChannel <- err
		}
	}(newListProduct)

	wg.Wait()
	close(errorChannel)

	for err := range errorChannel {
		if err != nil {
			return helpers.Response("Error goroutine", 500, err.Error())
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", len(newListProduct))
	response := helpers.Response(payload, 200, "Import json data to database is success")

	return response
}
