package model

import (
	"book-service/db"
	"github.com/google/uuid"
)

type Book struct {
	ID		string
	Title 	string
	Author 	string
	Price 	float64
}

func FetchAll() [] Book {
	db := db.ConnectDataBase()

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

	defer db.Close()

	return books
}

func CreateNewBook(title, author string, price float64) {
	db := db.ConnectDataBase()
	id := uuid.Must(uuid.NewRandom())
	sql := "INSERT INTO book(id, title, author, price) VALUES($1, $2, $3, $4);"
	insertion, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	insertion.Exec(id, title, author, price)
	defer db.Close()
}
