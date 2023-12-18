package models

import (
	"ecommerce_store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Categories struct {
	gorm.Model
	Category []Category `gorm:"many2many:category_product;"`
}

type Category struct {
	gorm.Model
	Name     string    `gorm:"column:category_name" json:"category_name"`
	Products []Product `gorm:"many2many:category_product;"`
}

type Product struct {
	gorm.Model
	ProductID   int      `gorm:"unique" json:"product_id"`
	Name        string   `gorm:"column:product_name" json:"product_name"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	Category    Category `gorm:"embedded"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Product{})
}
