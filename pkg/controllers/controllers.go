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
		if updateProduct.Price != 0 {
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

func GetCategoryById(w http.ResponseWriter, req *http.Request) {

}

func RemoveCategory(w http.ResponseWriter, req *http.Request) {

}

func UpdateCategory(w http.ResponseWriter, req *http.Request) {

}

func AllProducts(w http.ResponseWriter, req *http.Request) {

}

func ProductById(w http.ResponseWriter, req *http.Request) {

}
