package routes

import (
	"ecommerce_store/pkg/controllers"

	"github.com/gorilla/mux"
)

var AdminRoutes = func(router *mux.Router) {
	router.HandleFunc("/category", controllers.AddCategory).Methods("POST")
	router.HandleFunc("/category", controllers.AllCategories).Methods("GET")
	router.HandleFunc("/category/{category_id}", controllers.RemoveCategory).Methods("DELETE")
	router.HandleFunc("/category/{category_id}", controllers.GetCategoryById).Methods("GET")
	router.HandleFunc("/category/{category_id}", controllers.UpdateCategory).Methods("PUT")

	router.HandleFunc("/product", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/product", controllers.AllProducts).Methods("GET")
	router.HandleFunc("/product/{product_id}", controllers.ProductById).Methods("GET")
	router.HandleFunc("/product/{product_id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{product_id}", controllers.DeleteProduct).Methods("DELETE")
}
