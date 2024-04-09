package model

type Reports struct {
	Id           int    `gorm:"column:id"`
	ProductName  string `gorm:"column:name"`
	TotalQtySold int    `gorm:"column:total_sold"`
	BuyingPrice  int    `gorm:"column:buy"`
	SellingPrice int    `gorm:"column:sell"`
	GrossSales   int    `gorm:"column:gross_sales"`
	NettSales    int    `gorm:"column:nett_sales"`
}
