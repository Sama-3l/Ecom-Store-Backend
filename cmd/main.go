package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ecommerce_store/pkg/routes"

	"github.com/gorilla/mux"
)

func createPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := ":" + port

	return addr
}

func main() {
	r := mux.NewRouter()
	routes.AdminRoutes(r)
	http.Handle("/", r)
	port := createPort()
	fmt.Printf("Server online at %v\n", port)
	log.Fatal(http.ListenAndServe("localhost"+port, r))
}
