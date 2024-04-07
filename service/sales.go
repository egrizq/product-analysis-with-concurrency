package service

import (
	"fmt"
	"product-store-management/database"
	"product-store-management/helpers"
	"product-store-management/model"
	"strconv"
	"sync"
)

func formatSalesRecord(mapProductID map[string]int, salesRecord []string) *model.Sales {
	productId := mapProductID[salesRecord[0]]
	qty, _ := strconv.Atoi(salesRecord[1])
	date := helpers.ConvertDate(salesRecord[2])

	sales := &model.Sales{
		ProductId: productId,
		QtySold:   qty,
		Date:      date,
	}

	return sales
}

func ImportSalesToDatabase(csvRecords [][]string, mapProductID map[string]int) model.Response {
	var salesList []*model.Sales
	var wg sync.WaitGroup
	var salesListMutex sync.Mutex

	errorChannel := make(chan error)
	batchSize := 5000

	appendToSalesList := func(salesRecord []string) {
		defer wg.Done()

		sales := formatSalesRecord(mapProductID, salesRecord)
		salesListMutex.Lock()
		salesList = append(salesList, sales)
		salesListMutex.Unlock()
	}

	for index, salesRecord := range csvRecords {
		if index != 0 {
			wg.Add(1)

			go appendToSalesList(salesRecord)

			if len(salesList)%batchSize == 0 {
				wg.Wait()
			}
		}
	}
	wg.Wait()

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
				return
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
			return helpers.Response("Error channel", 200, err.Error())
		}
	}

	payload := fmt.Sprintf("Success insert %v rows of data into products table", totalSalesList)
	return helpers.Response(payload, 200, "Inserting CSV File into Sales Table")
}

func GetProductNameAndID() ([]model.ProductNameId, error) {
	var listProduct []model.ProductNameId

	query := "SELECT id, name FROM products;"
	if err := database.DB.Raw(query).Scan(&listProduct).Error; err != nil {
		return []model.ProductNameId{}, err
	}

	return listProduct, nil
}
