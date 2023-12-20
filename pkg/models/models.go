package models

import (
	"ecommerce_store/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Category struct {
	gorm.Model
	Name     string    `gorm:"column:category_name" json:"category_name"`
	Products []Product `gorm:"many2many:category_product;ForeignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	Name        string  `gorm:"column:product_name" json:"product_name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Category{}, &Product{})
}
