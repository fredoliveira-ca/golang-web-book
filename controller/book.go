package controller

import (
	"fmt"
	"log"
	"strconv"
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

func NewBook(w http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(w, "NewBook", nil)
}

func Insert(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	if req.Method == "POST" {
		fmt.Println(req.FormValue("title"))
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