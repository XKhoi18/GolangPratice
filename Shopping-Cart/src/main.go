package main

import (
	"shopping-cart/controllers/accountcontroller"
	"shopping-cart/controllers/cartcontroller"
	"shopping-cart/controllers/productcontroller"

	adminproductcontroller "shopping-cart/admin/controllers/productcontroller"

	"fmt"
	"net/http"
)

func main() {
	// Routing images and other assests
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//Home Page
	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)

	// Cart Page
	http.HandleFunc("/cart", cartcontroller.Index)
	http.HandleFunc("/cart/index", cartcontroller.Index)
	http.HandleFunc("/cart/buy", cartcontroller.Buy)
	http.HandleFunc("/cart/remove", cartcontroller.Remove)
	http.HandleFunc("/cart/order", cartcontroller.Order)

	// CRUD Product
	http.HandleFunc("/admin", adminproductcontroller.Index)
	http.HandleFunc("/productlist", adminproductcontroller.Index)
	http.HandleFunc("/productlist/add", adminproductcontroller.Add)
	http.HandleFunc("/productlist/addsubmit", adminproductcontroller.AddSubmit)
	http.HandleFunc("/productlist/delete", adminproductcontroller.Delete)
	http.HandleFunc("/productlist/edit", adminproductcontroller.Edit)
	http.HandleFunc("/productlist/editsubmit", adminproductcontroller.EditSubmit)

	//LogIn
	http.HandleFunc("/account", accountcontroller.Index)
	http.HandleFunc("/account/index", accountcontroller.Index)
	http.HandleFunc("/account/login", accountcontroller.Login)
	http.HandleFunc("/account/logout", accountcontroller.Logout)

	fmt.Println("Begin!")
	http.ListenAndServe(":3000", nil)

	fmt.Println("WebSite Started!")
}
