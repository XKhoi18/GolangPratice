package product_api

import (
	"Shop-Api/entities"
	"Shop-Api/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, err := productModel.FindAll()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(response, http.StatusOK, products)
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var productModel models.ProductModel
	products, err := productModel.Find(id)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(response, http.StatusOK, products)
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	var productModel models.ProductModel
	products, err := productModel.Search(keyword)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(response, http.StatusOK, products)
	}
}

func SearchPrices(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	min, _ := strconv.ParseFloat(vars["min"], 64)
	max, _ := strconv.ParseFloat(vars["max"], 64)
	var productModel models.ProductModel
	products, err := productModel.SearchPrices(min, max)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		respondWithJson(response, http.StatusOK, products)
	}
}

func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		var productModel models.ProductModel
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}

}

func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		var productModel models.ProductModel
		err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	var productModel models.ProductModel
	rowsAffected, err := productModel.Delete(id)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "err.Error()")
	} else {
		respondWithJson(response, http.StatusOK, map[string]int64{
			"RowsAffected": rowsAffected,
		})
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	// upload max 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory
	tempFile, err := ioutil.TempFile("images", handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of uploaded file into byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	//Enable CORS
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	w.WriteHeader(code)
	w.Write(response)
}
