package adminproductcontroller

import (
	"html/template"
	"net/http"
	"shopping-cart/entities"
	"shopping-cart/models"
	"strconv"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	//Get data from database
	products, _ := productModel.FindAll()
	sessions, _ := store.Get(request, "mysession")
	//Check Login status
	username := sessions.Values["username"]
	if username != "admin" {
		http.Redirect(response, request, "/account", http.StatusSeeOther)
	}
	isLogin := username == nil
	//Put data into []
	data := map[string]interface{}{
		"products": products,
		"username": username,
		"isLogin":  isLogin,
	}
	//Send data to UI
	tmp, _ := template.ParseFiles("admin/views/products/index.html")
	tmp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "mysession")
	username := sessions.Values["username"]
	if username != "admin" {
		http.Redirect(response, request, "/account", http.StatusSeeOther)
	}
	isLogin := username == nil
	data := map[string]interface{}{
		"username": username,
		"isLogin":  isLogin,
	}

	tmp, _ := template.ParseFiles("admin/views/products/addproduct.html")

	tmp.Execute(response, data)
}

func AddSubmit(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entities.Product
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Photo = request.Form.Get("photo")
	var productModel models.ProductModel
	productModel.Create(&product)

	http.Redirect(response, request, "/productlist", http.StatusSeeOther)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	//Get id from UI
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	//Call function
	var productModel models.ProductModel
	productModel.Delete(id)

	http.Redirect(response, request, "/productlist", http.StatusSeeOther)
}

func Edit(response http.ResponseWriter, request *http.Request) {
	//Get id from UI
	query := request.URL.Query()
	id, _ := strconv.ParseInt(query.Get("id"), 10, 64)
	//Call function
	var productModel models.ProductModel
	product, _ := productModel.Find(id)
	data := map[string]interface{}{
		"product": product,
	}

	tmp, _ := template.ParseFiles("admin/views/products/editproduct.html")
	tmp.Execute(response, data)
}

func EditSubmit(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var product entities.Product
	product.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
	product.Name = request.Form.Get("name")
	product.Price, _ = strconv.ParseFloat(request.Form.Get("price"), 64)
	product.Quantity, _ = strconv.ParseInt(request.Form.Get("quantity"), 10, 64)
	product.Photo = request.Form.Get("photo")
	var productModel models.ProductModel
	productModel.Update(product)

	http.Redirect(response, request, "/productlist", http.StatusSeeOther)
}
