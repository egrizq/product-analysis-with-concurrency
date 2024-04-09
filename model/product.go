package model

type Product struct {
	Id           int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Name         string `json:"name"`
	Stock        int    `json:"stock"`
	SellingPrice int    `json:"selling_price"`
	BuyingPrice  int    `json:"buying_price"`
}
