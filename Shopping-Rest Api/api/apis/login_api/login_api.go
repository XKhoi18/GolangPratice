package login_api

import (
	"Shop-Api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Login(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	username := vars["username"]
	password := vars["password"]
	var loginModel models.LoginModel
	logInStatus, role := loginModel.Login(username, password)
	if !logInStatus {
		respondWithError(response, http.StatusBadRequest, "Login Fail")
	} else {
		respondWithJson(response, http.StatusOK, map[string]string{
			"LoginStatus": "Login Success",
			"Role":        role,
		})
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
