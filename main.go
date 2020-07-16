package main

import (
	"fmt"
	"net/http"
)
type httpHandler struct {
	message string
}

func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	fmt.Fprint(resp, h.message)
}
func main(){

	h1:=httpHandler{message:"Conact"}

	http.Handle("/contact", h1)

	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "About")
	})
	http.HandleFunc("/contct", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Contact")
	})
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "index.html")
	})
	fmt.Println("Listening")
	http.ListenAndServe("localhost:8181", nil)
}
