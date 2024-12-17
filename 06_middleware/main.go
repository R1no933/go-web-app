package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/foo", FooHandler)

	handler := MyMiddleWare(mux)

	err := http.ListenAndServe(":3030", handler)
	if err != nil {
		panic(err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page.")
	w.Write([]byte("Home page."))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Foo page.")
	w.Write([]byte("Foo page."))
}

func MyMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before")
		h.ServeHTTP(w, r)
		fmt.Println("after")
	})
}
