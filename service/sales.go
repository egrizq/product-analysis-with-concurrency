package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
	"strconv"
	"sync"
)

func ImportSalesToDatabase(csvRecords [][]string) model.Response {
	var wg sync.WaitGroup

	errorChannel := make(chan error)
	batchSize := 5000

	salesList, err := insertSalesRecord(csvRecords, &wg)
	if err != nil {
		return helpers.Response("Error salesList", 404, err.Error())
	}
	totalSalesList := len(salesList)

	for i := 0; i < totalSalesList; i += batchSize {
		end := i + batchSize
		if end > totalSalesList {
			end = totalSalesList
		}

		wg.Add(1)
		batch := make([]*model.Sales, end-i)
		copy(batch, salesList[i:end])

		go func(sales []*model.Sales) {
			defer wg.Done()

			tx := database.DB.Begin()
			if tx.Error != nil {
				errorChannel <- tx.Error
			}

			if err := tx.Create(sales).Error; err != nil {
				tx.Rollback()
				errorChannel <- err
				return
			}

			if err := tx.Commit().Error; err != nil {
				errorChannel <- err
				return
			}
		}(batch)
	}

	wg.Wait()
	close(errorChannel)

	for err := range errorChannel {
		if err != nil {
			return helpers.Response("Error channel", 500, err.Error())
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", totalSalesList)
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}

func GetProductNameAndID() ([]model.Product, error) {
	var listProduct []model.Product

	// query := "SELECT id, name FROM products;"
	if err := database.DB.Select("id, name").Find(&listProduct).Error; err != nil {
		return []model.Product{}, err
	}

	return listProduct, nil
}

func MapProductID() (map[string]int, error) {
	listProductNameID, err := GetProductNameAndID()
	if err != nil {
		return map[string]int{}, err
	}

	// todo hashmap to build table relation with product.id from sales table
	productID := make(map[string]int)
	for _, product := range listProductNameID {
		productID[product.Name] = product.Id
	}

	return productID, nil
}

func formatSalesRecord(productID map[string]int, salesRecord []string) *model.Sales {
	id := productID[salesRecord[0]]
	qty, _ := strconv.Atoi(salesRecord[1])
	date := helpers.ConvertDate(salesRecord[2])

	sales := &model.Sales{
		ProductId: id,
		QtySold:   qty,
		Date:      date,
	}

	return sales
}

func insertSalesRecord(csvRecords [][]string, wg *sync.WaitGroup) ([]*model.Sales, error) {
	var salesList []*model.Sales
	var salesListMutex sync.Mutex

	batch := 5000

	productID, err := MapProductID()
	if err != nil {
		return []*model.Sales{}, err
	}

	appendToSalesList := func(salesRecord []string) {
		defer wg.Done()

		sales := formatSalesRecord(productID, salesRecord)
		salesListMutex.Lock()
		salesList = append(salesList, sales)
		salesListMutex.Unlock()
	}

	for index, salesRecord := range csvRecords {
		if index != 0 {
			wg.Add(1)
			go appendToSalesList(salesRecord)

			if index%batch == 0 {
				wg.Wait()
			}
		}
	}
	wg.Wait()

	return salesList, nil
}
