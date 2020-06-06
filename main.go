package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
  )

type Book struct {
	Title string
	Author string
	Price float64
}  

var templates = template.Must(template.ParseGlob("templates/*.html")) 

func main() {
	http.HandleFunc("/", fetchData)
	http.ListenAndServe(":8080", nil)
}

func fetchData(w http.ResponseWriter, req *http.Request) {
	db := connectDataBase()

	records, err := db.Query("select * from book")

	if err != nil {
		panic(err)
	}

	book := Book{}
	books := [] Book{}

	for records.Next() {
		var id, title, author string
		var price float64

		err = records.Scan(&id, &title, &author, &price)
		if err != nil {
			panic(err.Error())
		}

		book.Title = title
		book.Author = author
		book.Price = price

		books = append(books, book)
	}

	templates.ExecuteTemplate(w, "Index", books)

	defer db.Close()
}

func connectDataBase() *sql.DB {
	connection := fmt.Sprintf("host=%s port=%d user=%s  password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}