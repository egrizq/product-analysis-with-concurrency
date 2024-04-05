package model

type Sales struct {
	Id        int `gorm:"primaryKey;AUTO_INCREMENT"`
	ProductId int `json:"product_id"`
	QtySold   int `json:"qty_sold"`
}
