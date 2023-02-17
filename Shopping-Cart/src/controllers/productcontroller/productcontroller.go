package productcontroller

import (
	"html/template"
	"net/http"
	"shopping-cart/models"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	//Get data from database
	products, _ := productModel.FindAll()
	//Get username
	sessions, _ := store.Get(request, "mysession")
	username := sessions.Values["username"]
	isLogin := username == nil
	//Put data into an interface
	data := map[string]interface{}{
		"products": products,
		"username": username,
		"isLogin":  isLogin,
	}
	//Send data to UI
	tmp, _ := template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)
}
