package main

import (
	"fmt"
	"net/http"
)

func main(){

	fs:=http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "static/about.html")
	})
	http.HandleFunc("/contact", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Contatc")
	})
	http.ListenAndServe("localhost:8181", nil)
}
