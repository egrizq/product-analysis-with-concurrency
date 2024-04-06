package helpers

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"product-store-management/model"
)

func ReadCSV(file string) ([][]string, error) {
	fileCSV, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer fileCSV.Close()

	reader := csv.NewReader(fileCSV)
	return reader.ReadAll()
}

func ReadJSON(file string) ([]model.Product, error) {
	fileJson, err := os.ReadFile(file)
	if err != nil {
		return []model.Product{}, err
	}

	var listProduct []model.Product
	err = json.Unmarshal(fileJson, &listProduct)
	if err != nil {
		return []model.Product{}, err
	}

	return listProduct, nil
}
