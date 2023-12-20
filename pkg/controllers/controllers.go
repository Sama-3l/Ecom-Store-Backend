package controllers

import (
	"ecommerce_store/pkg/models"
	"ecommerce_store/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddCategory(w http.ResponseWriter, req *http.Request) {
	Category := &models.Category{}
	utils.ParseBody(req, Category)
	res, _ := json.Marshal(Category)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println(Category)
}

func AllCategories(w http.ResponseWriter, req *http.Request) {

}

func GetCategoryById(w http.ResponseWriter, req *http.Request) {

}

func RemoveCategory(w http.ResponseWriter, req *http.Request) {

}

func UpdateCategory(w http.ResponseWriter, req *http.Request) {

}

func AddProduct(w http.ResponseWriter, req *http.Request) {

}

func AllProducts(w http.ResponseWriter, req *http.Request) {

}

func ProductById(w http.ResponseWriter, req *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, req *http.Request) {

}

func UpdateProduct(w http.ResponseWriter, req *http.Request) {

}
