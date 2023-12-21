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
	c := Category.AddCategory()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	fmt.Println(c)
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
