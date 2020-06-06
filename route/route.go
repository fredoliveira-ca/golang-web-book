package route

import (
	"net/http"
	"book-service/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.FetchData)
	http.HandleFunc("/new-book", controller.NewBook)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
}