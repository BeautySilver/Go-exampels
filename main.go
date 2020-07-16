package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func productsHandler( writer http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	id:=vars["id"]
	cat:= vars["cat"]
	response:=fmt.Sprintf("Product category=%s id=%s", cat, id)
	fmt.Fprint(writer, response)
}

func indexHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Index")
}
func main(){

	router:= mux.NewRouter()
    router.HandleFunc("/products/category/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/articles/{id:[0-9]}", productsHandler)
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)
    http.ListenAndServe("localhost:8181", nil)
}
