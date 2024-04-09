package service

import (
	"log"
	"product-store-management/database"
	"product-store-management/model"

	"gorm.io/gorm"
)

/*
	select product -> id & stock
	store product

	select sales -> product_id & SUM(qty_sold)
	groupby sales -> product_id
	store sales

	calculate product stock - qty_sold
	where product id = sales product id
	update product

*/

type Query interface {
	GetProducts() []model.Product
}

func Reports(db *gorm.DB) Query {
	return &Begin{
		db: db,
	}
}

type Begin struct {
	db *gorm.DB
}

func SalesReports() []model.Reports {

	var salesReports []model.Reports

	query := `
		SELECT 
			products.id,
			products.name,
			SUM(sales.qty_sold) as total_sold,
			products.buying_price as buy,
			products.selling_price as sell,
			products.selling_price * sum(sales.qty_sold) as gross_sales,
			(products.selling_price * sum(sales.qty_sold)) - (products.buying_price * sum(sales.qty_sold)) as nett_sales
		FROM products
		JOIN sales
		ON products.id = sales.product_id
		GROUP BY sales.product_id, products.id
		ORDER BY products.id ASC;		
	`
	err := database.DB.Raw(query).Scan(&salesReports).Error

	if err != nil {
		log.Fatal(err.Error())
	}

	return salesReports
}

func (begin *Begin) GetProducts() []model.Product {
	var products []model.Product

	if err := begin.db.Find(&products).Error; err != nil {
		log.Fatal(err)
	}

	return products
}
