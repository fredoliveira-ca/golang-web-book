package controller

import (
	"net/http"
	"html/template"
	"book-service/model"
	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html")) 

func FetchData(w http.ResponseWriter, req *http.Request) {
	books := model.FetchAll()
	templates.ExecuteTemplate(w, "Index", books)
}