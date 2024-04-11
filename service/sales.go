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

	salesList, err := insertSalesRecord(csvRecords, &wg)
	if err != nil {
		return helpers.Response("Error salesList", 404, err.Error())
	}

	rangeSalesList := len(salesList)
	batchSize := 5000
	errorChannel := make(chan error, rangeSalesList/batchSize+1)

	// todo looping
	for i := 0; i < rangeSalesList; i += batchSize {
		end := i + batchSize
		if end > rangeSalesList {
			end = rangeSalesList
		}

		wg.Add(1)
		salesBatch := make([]*model.Sales, end-i)
		copy(salesBatch, salesList[i:end])

		go func(salesBatch []*model.Sales) {
			defer wg.Done()

			if err := SaveBatch(salesBatch); err != nil {
				errorChannel <- err
				return
			}
		}(salesBatch)

	}
	wg.Wait()
	close(errorChannel)

	var rollbackErr error
	for err := range errorChannel {
		if err != nil {
			return helpers.Response("Error occurred during processing, transaction rolled back", 500, rollbackErr.Error())
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", rangeSalesList)
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}

func SaveBatch(salesBatch []*model.Sales) error {
	tx := database.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(salesBatch).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func GetProductNameAndID() ([]model.Product, error) {
	var listProduct []model.Product

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
