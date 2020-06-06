package route

import (
	"net/http"
	"book-service/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.FetchData)
}