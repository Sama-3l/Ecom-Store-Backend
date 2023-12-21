package models

import (
	"ecommerce_store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Category struct {
	gorm.Model
	CategoryName string `gorm:"unique" json:"category_name"`
	Products     []Product
}

type Product struct {
	gorm.Model
	ProdID      string `gorm:"unique" json:"prod_id"`
	ProductName string `json:"product_name"`
	Price       float32
	Description string
	CategoryID  uint
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Category{}, &Product{})
}

func (category *Category) AddCategory() (*Category, error) {
	db.NewRecord(&category)
	err := db.Create(&category).Error
	return category, err
}

func GetProductByCategory(id int64) []Product {
	var products []Product
	db.Where("category_id = ?", id).Find(&products)
	return products
}
