package cartcontroller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"shopping-cart/entities"
	"shopping-cart/models"
	"strconv"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	//Create variable for session and cart items
	sessions, _ := store.Get(request, "mysession")
	username := sessions.Values["username"]
	//Check login status
	if username != "user" || username == nil {
		http.Redirect(response, request, "/account", http.StatusSeeOther)
	}

	var cart []entities.Item
	var data map[string]interface{}
	//json.Unmarshal([]byte(strCart), &cart)
	if sessions.Values["cart"] != nil {
		strCart, _ := sessions.Values["cart"].(string)
		if err2 := json.Unmarshal([]byte(strCart), &cart); err2 != nil {
			fmt.Println(err2)
			os.Exit(0)
		}
		// Put data into interface
		// data ["cart"] = cart
		// data ["total"] = total(cart)
		data = map[string]interface{}{
			"cart":     cart,
			"total":    total(cart),
			"username": username,
		}
	}
	// Call function to calculate total for each item
	tmp, _ := template.New("index.html").Funcs(template.FuncMap{
		"totalItem": func(item entities.Item) float64 {
			return item.Product.Price * float64(item.Quantity)
		},
	}).ParseFiles("views/cart/index.html")
	//Send data to UI
	tmp.Execute(response, data)
}

func Buy(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)

	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	//Call current session
	sessions, _ := store.Get(request, "mysession")
	cart := sessions.Values["cart"]
	// Check if cart have any item, nil then add new
	if cart == nil {
		var cart []entities.Item
		cart = append(cart, entities.Item{
			Product:  product,
			Quantity: 1,
		})
		bytesCart, _ := json.Marshal(cart)
		sessions.Values["cart"] = string(bytesCart)
		sessions.Save(request, response)
		// Check if cart have the item, if no then add new otherwise Quantity++
	} else {
		strCart, _ := sessions.Values["cart"].(string)
		var cart []entities.Item
		if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		index := exists(id, cart)

		if index == -1 {
			cart = append(cart, entities.Item{
				Product:  product,
				Quantity: 1,
			})
		} else {
			cart[index].Quantity++
		}
		bytesCart, _ := json.Marshal(cart)
		sessions.Values["cart"] = string(bytesCart)
	}
	//Save session
	sessions.Save(request, response)
	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}

func Remove(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	//Call current session
	sessions, _ := store.Get(request, "mysession")

	strCart, _ := sessions.Values["cart"].(string)
	var cart []entities.Item
	if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// Remove the item
	index := exists(id, cart)
	cart = remove(cart, index)
	// Save new session after remove
	bytesCart, _ := json.Marshal(cart)
	sessions.Values["cart"] = string(bytesCart)
	sessions.Save(request, response)

	http.Redirect(response, request, "/cart", http.StatusSeeOther)
}

func Order(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "mysession")
	strCart, _ := sessions.Values["cart"].(string)
	username := sessions.Values["username"].(string)
	var cart []entities.Item
	if err := json.Unmarshal([]byte(strCart), &cart); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	var cartModel models.CartModel
	cartModel.Order(cart, username)

	delete(sessions.Values, "cart")
	sessions.Save(request, response)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}

func exists(id int64, cart []entities.Item) int {
	for i := 0; i < len(cart); i++ {
		if cart[i].Product.Id == id {
			return i
		}
	}
	return -1
}

func total(cart []entities.Item) float64 {
	var s float64 = 0
	for _, item := range cart {
		s += item.Product.Price * float64(item.Quantity)
	}
	return s
}

func remove(cart []entities.Item, index int) []entities.Item { // []entities.Item is return type
	//Shift cart[i+1:] left one index.
	//copy (cart[index:], cart[index+1:])
	cart = append(cart[0:index], cart[index+1:]...)
	//Truncate slice
	//return cart[:len(cart)-1]
	return cart
}
