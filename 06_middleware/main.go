package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/foo", FooHandler)

	mdwrs := []func(http.Handler) http.Handler{
		LoggingMiddleWare,
		SecondMiddleWare,
	}

	h := http.Handler(mux)
	for i := len(mdwrs) - 1; i >= 0; i-- {
		h = mdwrs[i](h)
	}

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home page.\n"))
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Foo page.\n"))
}

func LoggingMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s\n", r.RequestURI)
	})
}

func SecondMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
