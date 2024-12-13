package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/form", FormHandler)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		panic(err)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		io.WriteString(w, "405 Method not allowed!")
		return
	}

	first := r.FormValue("first")

	io.WriteString(w, "OK: "+first)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	file, header, err := r.FormFile("test-file")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	fmt.Println(header.Filename, header.Size)

	ext := path.Ext(header.Filename)
	tmpFile, err := os.CreateTemp("/tmp", "*"+ext)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}
	defer tmpFile.Close()

	fmt.Println(tmpFile.Name())

	bts, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	_, err = tmpFile.Write(bts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}

	fmt.Println(w, "OK! File uploaded succesfull.")
}
