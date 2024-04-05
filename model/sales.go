package model

import "time"

type Sales struct {
	Id          int       `gorm:"primaryKey"`
	ProductName string    `gorm:"column:product_name"`
	Qty         int       `gorm:"column:qty_sold"`
	Date        time.Time `gorm:"column:sale_at"`
}
