package models

import (
	"ecommerce_store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// type Category struct {
// 	gorm.Model
// 	Name     string    `gorm:"column:category_name" json:"category_name"`
// 	Products []Product `gorm:"many2many:category_product;ForeignKey:CategoryID" json:"products"`
// }

// type Product struct {
// 	gorm.Model
// 	Name        string   `gorm:"column:product_name" json:"product_name"`
// 	Price       float32  `json:"price"`
// 	Description string   `json:"description"`
// 	CategoryID  uint     `json:"category_id"`
// 	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
// }

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name"`
	Products     []Product
}

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Price       float32
	Description string
	CategoryID  uint
	// Category     Category `gorm:"foreignKey:CategoryID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Category{}, &Product{})
}

func (category *Category) AddCategory() *Category {
	// db.NewRecord(&category)
	db.Create(&category)
	return category
}

func GetProductByCategory(id int64) []Product {
	var products []Product
	db.Where("category_id = ?", id).Find(&products)
	return products
}
