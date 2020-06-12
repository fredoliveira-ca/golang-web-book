package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/fredoliveira-ca/book-service/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func FetchData(w http.ResponseWriter, req *http.Request) {
	books := model.FetchAll()
	templates.ExecuteTemplate(w, "Index", books)
}

func Create(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "NewBook", nil)
}

func Insert(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		title := req.FormValue("title")
		author := req.FormValue("author")
		price, err := strconv.ParseFloat(req.FormValue("price"), 64)

		if err != nil {
			log.Println("Error while trying parse price value.")
		}

		model.CreateNewBook(title, author, price)

	}

	http.Redirect(w, req, "/", 301)
}

func Delete(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	model.DeleteBook(id)

	http.Redirect(w, req, "/", 301)
}

func Edit(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	book := model.FindOne(id)
	templates.ExecuteTemplate(w, "EditBook", book)
}

func Update(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		id := req.FormValue("id")
		title := req.FormValue("title")
		author := req.FormValue("author")
		price, err := strconv.ParseFloat(req.FormValue("price"), 64)

		if err != nil {
			log.Println("Error while trying parse price value.")
		}

		model.UpdateBook(id, title, author, price)
	}

	http.Redirect(w, req, "/", 301)
}
