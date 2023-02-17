package handler

import (
	"encoding/json"
	"net/http"
	"rest-api/data"
	"rest-api/entities"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProduct(writer http.ResponseWriter, request *http.Request) {
	responseWithJson(writer, http.StatusOK, data.Products)
}

func GetProductById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid product id"})
		return
	}

	for _, product := range data.Products {
		if product.ID == id {
			responseWithJson(writer, http.StatusOK, product)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var newProduct entities.Product
	if err := json.NewDecoder(request.Body).Decode(&newProduct); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid JSON body"})
		return
	}

	newProduct.ID = generateId(data.Products)
	data.Products = append(data.Products, newProduct)

	responseWithJson(writer, http.StatusCreated, newProduct)
}

func UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Product id"})
		return
	}

	var updateProduct entities.Product
	if err := json.NewDecoder(request.Body).Decode(&updateProduct); err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Json body"})
		return
	}
	updateProduct.ID = id

	for i, product := range data.Products {
		if product.ID == id {
			data.Products[i] = updateProduct
			responseWithJson(writer, http.StatusOK, updateProduct)
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJson(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Product id"})
		return
	}

	for i, product := range data.Products {
		if product.ID == id {
			data.Products = append(data.Products[:i], data.Products[i+1:]...)
			responseWithJson(writer, http.StatusOK, map[string]string{"message": "Product was deleted"})
			return
		}
	}

	responseWithJson(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}

func generateId(products []entities.Product) int {
	var maxId int
	for _, product := range products {
		if product.ID > maxId {
			maxId = product.ID
		}
	}

	return maxId + 1
}
