package main

import (
	"fmt"
	"net/http"
	"uploadFile/apis/upload_api"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/upload", upload_api.UpLoadFile).Methods("POST")
	fmt.Println("Begin !!!")
	err := http.ListenAndServe(":5001", router)
	if err != nil {
		fmt.Println(err)
	}
}
