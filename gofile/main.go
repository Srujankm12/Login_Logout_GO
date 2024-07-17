package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var filename = "login.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("the error is ", err)
	}

	type srujan struct {
		Name  string
		Level int
	}
	err = t.ExecuteTemplate(w, filename, srujan{" log in", 10})
	if err != nil {
		fmt.Println("the error while execution ", err)
	}

}

var userDB = map[string]string{
	"user1": "pass1",
}

func loginsubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if userDB[username] == password {

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "you are logged in")

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "you are not logged in")
	}

}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/loginsubmit":
		loginsubmit(w, r)
	default:
		fmt.Fprintf(w, "Hello World")
	}

}

func main() {

	http.HandleFunc("/", Handler)
	fmt.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
	// http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)

}
