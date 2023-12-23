package controllers

import (
	"ecommerce_store/pkg/models"
	"ecommerce_store/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCategory(w http.ResponseWriter, req *http.Request) {
	Category := &models.Category{}
	utils.ParseBody(req, Category)
	c, err := Category.AddCategory()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		res, _ := json.Marshal(c)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
		fmt.Println(c)
	}

}

func GetAllCategories(w http.ResponseWriter, req *http.Request) {
	c, err := models.GetAllCategories()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		res, _ := json.Marshal(c)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetCategoryById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_string := vars["category_id"]
	id, err := strconv.ParseInt(id_string, 10, 12)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	} else {
		category, db := models.GetCategoryById(id)
		if db.Error != nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			res, _ := json.Marshal(category)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	}
}

func RemoveCategory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id_string := vars["category_id"]
	rec := vars["record"]
	fmt.Println(id_string)
	maintain_rec, _ := strconv.ParseBool(rec)
	id, _ := strconv.ParseInt(id_string, 10, 12)
	category, e := models.RemoveCategory(id, maintain_rec)
	if e != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		res, _ := json.Marshal(category)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func UpdateCategory(w http.ResponseWriter, req *http.Request) {
	updateCategory := &models.Category{}
	utils.ParseBody(req, updateCategory)
	vars := mux.Vars(req)
	category_id := vars["category_id"]
	id, _ := strconv.ParseInt(category_id, 10, 12)
	category, db := models.GetCategoryById(id)
	if db.Error != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		if updateCategory.CategoryName != "" {
			category.CategoryName = updateCategory.CategoryName
		}
		if updateCategory.Products != nil {
			category.Products = updateCategory.Products
		}
		fmt.Println(category)
		err := db.Save(&category).Error
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			res, _ := json.Marshal(category)
			w.Write(res)
		} else {
			res, _ := json.Marshal(category)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}

	}
}

func GetProductByCategory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	categoryID, _ := strconv.ParseInt(vars["category_id"], 10, 12)
	products := models.GetProductByCategory(categoryID)
	res, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println(products)
}

func AddProductToCategory(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	category := vars["category_name"]
	products := &models.RequestBody{}
	utils.ParseBody(req, products)
	currentCategory, err := models.AddProductToCategory(category, *products)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		fmt.Println(currentCategory)
		res, _ := json.Marshal(currentCategory)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func GetProductById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	prod_id := vars["product_id"]
	prod, err := models.GetProductById(prod_id)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		res, _ := json.Marshal(prod)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func UpdateProduct(w http.ResponseWriter, req *http.Request) {
	updateProduct := &models.Product{}
	utils.ParseBody(req, updateProduct)
	vars := mux.Vars(req)
	prod_id := vars["product_id"]
	prod, db := models.GetProductById(prod_id)
	if db.Error != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		if updateProduct.Description != "" {
			prod.Description = updateProduct.Description
		}
		if updateProduct.ProductName != "" {
			prod.ProductName = updateProduct.ProductName
		}
		if updateProduct.Price != "" {
			prod.Price = updateProduct.Price
		}
		if updateProduct.ProdID != "" {
			prod.ProdID = updateProduct.ProdID
		}
		err := db.Save(&prod).Error
		if err != nil {
			w.WriteHeader(http.StatusConflict)
		} else {
			res, _ := json.Marshal(prod)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}

	}
}

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	prod_id := vars["product_id"]
	prod, err := models.DeleteProduct(prod_id)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		res, _ := json.Marshal(prod)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func AllProducts(w http.ResponseWriter, req *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
	} else {
		res, _ := json.Marshal(products.Products)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func ProductById(w http.ResponseWriter, req *http.Request) {

}
