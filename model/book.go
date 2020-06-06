package model

import "book-service/db"

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

