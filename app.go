package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/fredoliveira-ca/book-service/route"
)

func main() {
	route.LoadRoutes()
	http.ListenAndServe(":8010", nil)
}
