package main

import (
	"Shop-Api/apis/cart_api"
	"Shop-Api/apis/login_api"
	"Shop-Api/apis/product_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/find/{id}", product_api.Find).Methods("GET")
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods("GET")
	router.HandleFunc("/api/product/search/{min}/{max}", product_api.SearchPrices).Methods("GET")
	router.HandleFunc("/api/product/create", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product/update", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/product/delete/{id}", product_api.Delete).Methods("DELETE")
	//router.HandleFunc("/api/product/upload", product_api.Update).Methods("POST")

	router.HandleFunc("/api/cart/order", cart_api.Order).Methods("POST")

	router.HandleFunc("/api/login/{username}/{password}", login_api.Login).Methods("GET")

	//Enable CORS
	handler := cors.Default().Handler(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		//AllowedHeaders:     []string{"Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		//OptionsPassthrough: true,
	})
	handler = c.Handler(handler)

	fmt.Println("Begin!!!")
	err := http.ListenAndServe(":3001", handler)
	if err != nil {
		fmt.Println(err)
	}
}
