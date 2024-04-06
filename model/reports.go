package model

import "time"

type Reports struct {
	Id           int       `json:"id" gorm:"PrimaryKey;AUTO_INCEREMENT"`
	ProductId    int       `json:"product_id"`
	BuyingPrice  int       `json:"buying_price"`
	SellingPrice int       `json:"selling_price"`
	TotalQtySold int       `json:"total_qty_sold"`
	Nett         int       `json:"nett_sale_product"`
	Gross        int       `json:"gross_sale_product"`
	Year         time.Time `json:"year"`
}
