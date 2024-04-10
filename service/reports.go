package service

import (
	"fmt"
	"log"
	"product-store-management/helpers"
	"product-store-management/model"

	"gorm.io/gorm"
)

type Process interface {
	Analysis() []model.Reports
	Insert(reports []model.Reports) model.Response
	UpdateProducts() model.Response
}

func ReportsInit(db *gorm.DB) Process {
	return &Begin{
		db: db,
	}
}

type Begin struct {
	db *gorm.DB
}

func (begin *Begin) Analysis() []model.Reports {
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
	if err := begin.db.Raw(query).Scan(&salesReports).Error; err != nil {
		log.Fatal(err.Error())
	}

	return salesReports
}

func (begin *Begin) Insert(reports []model.Reports) model.Response {
	if err := begin.db.Create(&reports).Error; err != nil {
		return helpers.Response("Error inserting reports to database", 500, err.Error())
	}

	updateProducts := begin.UpdateProducts()
	if updateProducts.StatusCode == 500 {
		return updateProducts
	}

	payload := fmt.Sprintf("Success analysis and processing %v rows of data into reports table", len(reports))
	return helpers.Response(payload, 200, "Reports has been created")
}

func (begin *Begin) UpdateProducts() model.Response {
	query := `
		UPDATE products
		SET stock = stock - sales_summary.total_sold
		FROM (
			SELECT product_id, SUM(qty_sold) AS total_sold
			FROM sales
			GROUP BY product_id
		) as sales_summary
		WHERE id = sales_summary.product_id;
	`

	if err := begin.db.Exec(query).Error; err != nil {
		return helpers.Response("Error update product to database", 500, err.Error())
	}

	return model.Response{} // return nil if it the query was ok
}
