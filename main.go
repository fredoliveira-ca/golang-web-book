package main

import (
	"net/http"
	"book-service/route"
	_ "github.com/lib/pq"
)

func main() {	
	route.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}


