package main

import (
	"fmt"
	"net/http"
)

func main(){

http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "form.html")
})
http.HandleFunc("/postform", func(writer http.ResponseWriter, request *http.Request) {
	name:= request.FormValue("username")
	age:=request.FormValue("userage")

	fmt.Fprintf(writer,"Name: %s Age: %s", name, age)
})
http.ListenAndServe("localhost:8181", nil)
}
