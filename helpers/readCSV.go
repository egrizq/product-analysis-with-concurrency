package helpers

import (
	"encoding/csv"
	"log"
	"os"
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
