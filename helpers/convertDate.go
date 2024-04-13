package helpers

import (
	"log"
	"time"
)

func ConvertDate(dateStr string) time.Time {
	dateFormat := "02/01/2006"
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	return date
}
