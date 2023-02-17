package cart_api

import (
	"Shop-Api/entities"
	"Shop-Api/models"
	"encoding/json"
	"net/http"
)

func Order(response http.ResponseWriter, request *http.Request) {
	var cart []entities.Product
	err := json.NewDecoder(request.Body).Decode(&cart)
	//err := json.NewDecoder(request.Body).Decode(&product)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		var cartModel models.CartModel
		err2 := cartModel.Order(cart, "user")
		if !err2 {
			respondWithError(response, http.StatusBadRequest, "Error !!!")
		} else {
			respondWithJson(response, http.StatusOK, "Success")
		}
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

// func Index(response http.ResponseWriter, request *http.Request) {
// 	//Create variable for session and cart items
// 	sessions, _ := store.Get(request, "mysession")
// 	//username := sessions.Values["username"]
// 	//Check login status
// 	// if username != "user" || username == nil {
// 	// 	http.Redirect(response, request, "/account", http.StatusSeeOther)
// 	// }

// 	var cart []entities.Item
// 	if sessions.Values["cart"] != nil {
// 		strCart, _ := sessions.Values["cart"].(string)
// 		if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
// 			respondWithError(response, http.StatusBadRequest, err.Error())
// 		} else {
// 			respondWithJson(response, http.StatusOK, map[string]interface{}{
// 				"cart":  cart,
// 				"total": total(cart),
// 			})
// 		}
// 	} else {
// 		respondWithJson(response, http.StatusOK, map[string]interface{}{
// 			"cart":  cart,
// 			"total": total(cart),
// 		})
// 	}
// }

// func Buy(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	id, _ := strconv.ParseInt(vars["id"], 10, 64)

// 	var productModel models.ProductModel
// 	product, _ := productModel.Find(id)
// 	//Call current session
// 	sessions, _ := store.Get(request, "mysession")
// 	carts := sessions.Values["cart"]
// 	var cart []entities.Item
// 	// Check if cart have any item, nil then add new
// 	if carts == nil {
// 		carts = append(cart, entities.Item{
// 			Product:  product,
// 			Quantity: 1,
// 		})
// 		bytesCart, _ := json.Marshal(cart)
// 		sessions.Values["cart"] = string(bytesCart)
// 		sessions.Save(request, response)
// 		// Check if cart have the item, if no then add new otherwise Quantity++
// 	} else {
// 		strCart, _ := sessions.Values["cart"].(string)
// 		if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
// 			respondWithError(response, http.StatusBadRequest, err.Error())
// 		}

// 		index := exists(id, cart)
// 		if index == -1 {
// 			carts = append(cart, entities.Item{
// 				Product:  product,
// 				Quantity: 1,
// 			})
// 		} else {
// 			cart[index].Quantity++
// 		}
// 		bytesCart, _ := json.Marshal(cart)
// 		sessions.Values["cart"] = string(bytesCart)
// 	}
// 	//Save session
// 	sessions.Save(request, response)
// 	// respondWithJson(response, http.StatusOK, "Product "+strconv.Itoa(int(id))+" Added")
// 	respondWithJson(response, http.StatusOK, cart)
// }

// func Remove(response http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	id, _ := strconv.ParseInt(vars["id"], 10, 64)
// 	//Call current session
// 	sessions, _ := store.Get(request, "mysession")

// 	strCart, _ := sessions.Values["cart"].(string)
// 	var cart []entities.Item
// 	if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
// 		respondWithError(response, http.StatusBadRequest, err.Error())
// 	}
// 	// Remove the item
// 	index := exists(id, cart)
// 	if index == -1 {
// 		respondWithError(response, http.StatusBadRequest, "Can't find ID")
// 	} else {
// 		cart = remove(cart, index)
// 		// Save new session after remove
// 		bytesCart, _ := json.Marshal(cart)
// 		sessions.Values["cart"] = string(bytesCart)
// 		sessions.Save(request, response)
// 		respondWithJson(response, http.StatusOK, "Product "+strconv.Itoa(int(id))+" Deleted")
// 	}
// }
