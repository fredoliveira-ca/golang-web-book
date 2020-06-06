package main

import (
	"net/http"
	"html/template"
	"book-service/model"
	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html")) 

func main() {
	http.HandleFunc("/", fetchData)
	http.ListenAndServe(":8080", nil)
}

func fetchData(w http.ResponseWriter, req *http.Request) {
	books := model.FetchAll()
	templates.ExecuteTemplate(w, "Index", books)
}

