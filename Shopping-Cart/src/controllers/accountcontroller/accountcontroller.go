package accountcontroller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("views/account/index.html")
	tmp.Execute(response, nil)
}

func Login(response http.ResponseWriter, request *http.Request) {
	//Get input from UI
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")
	if username == "user" && password == "123" {
		sessions, _ := store.Get(request, "mysession")
		sessions.Values["username"] = username
		sessions.Save(request, response)
		http.Redirect(response, request, "/product", http.StatusSeeOther)
	} else if username == "admin" && password == "123" {
		sessions, _ := store.Get(request, "mysession")
		sessions.Values["username"] = username
		sessions.Save(request, response)
		http.Redirect(response, request, "/productlist", http.StatusSeeOther)
	} else {
		data := map[string]interface{}{
			"errLogin": true,
		}
		tmp, _ := template.ParseFiles("views/account/index.html")
		tmp.Execute(response, data)
	}
}

func Logout(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "mysession")
	//remove session on logout
	delete(sessions.Values, "username")
	delete(sessions.Values, "cart")
	sessions.Save(request, response)
	http.Redirect(response, request, "/product", http.StatusSeeOther)
}
