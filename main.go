package main

import (
	"html/template"
	"net/http"
)
type ViewData struct {
	Title string
	Message string
}
func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := ViewData{
			Title:   "World Cup",
			Message: "Fifa",
		}

		tmpl := template.Must(template.New("data").Parse(`<div> 
    <h1> {{ .Title}}</h1>
     <p>{{ .Message}}</p>
      </div>`))
		tmpl.Execute(writer, data)
	})
	http.ListenAndServe("localhost:8181", nil)
}
