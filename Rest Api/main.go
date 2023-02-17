package main

import (
	"log"
	"net/http"
	"rest-api/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/product", handler.GetAllProduct).Methods(http.MethodGet)
	r.HandleFunc("/api/product/{id}", handler.GetProductById).Methods(http.MethodGet)
	r.HandleFunc("/api/product", handler.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/product/{id}", handler.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/api/product/{id}", handler.DeleteProduct).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}
