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

type RequestBody struct {
	Products []Product `json:"products"`
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

func GetAllCategories() ([]Category, error) {
	var Categories []Category
	err := db.Preload("Products").Find(&Categories).Error
	return Categories, err
}

func GetProductByCategory(id int64) []Product {
	var products []Product
	db.Where("category_id = ?", id).Find(&products)
	return products
}

func AddProductToCategory(category string, products RequestBody) (*Category, error) {
	var currentCategory Category
	if err := db.Where("category_name = ?", category).First(&currentCategory).Error; err != nil {
		return &currentCategory, err
	}
	var err error
	for i := range products.Products {
		products.Products[i].CategoryID = currentCategory.ID
		err = db.Create(&products.Products[i]).Error
		if err != nil {
			return &currentCategory, err
		}
	}
	err = db.Where("category_name = ?", category).Preload("Products").Find(&currentCategory).Error
	return &currentCategory, err

	// if err := db.Create(&products.Products).Error; err != nil {
	// 	return &currentCategory, err
	// }

	// db := db.Where("category_name = ?", category).Preload("Products").Find(&currentCategory)
	// return &currentCategory, db.Error
}
