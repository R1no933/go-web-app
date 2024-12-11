package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/hello/{username}", userPage)
	r.HandleFunc(`/product/{id:\d+}`, productPage)
	r.HandleFunc(`/form`, formPage).Methods("POST", "PUT")
	r.NotFoundHandler = http.HandlerFunc(handler404)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}

func userPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	fmt.Fprintf(w, "User page. Username = %s", username)
}

func productPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Product id = %s", id)
}

func formPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Form page")
}

func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Error 404! Page not found!")
}
