package main

import (
	"net/http"
	"html/template"
)

type Book struct {
	Title string
	Author string
	Price float64
}  

var templates = template.Must(template.ParseGlob("templates/*.html")) 

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	books := [] Book {
		{ Title: "SOLID", Author: "Uncle Bob", Price: 23.2 },
		{"Clean Code", "Uncle Bob", 32.3},
		{"Building Microservices: Designing Fine-Grained Systems", "Sam Newma", 32.3},
		{"Structure and Interpretation of Computer Programs", "Gerald Jay Sussman and Hal Abelson", 32.3},
	}

	templates.ExecuteTemplate(w, "Index", books)
}