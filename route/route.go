package route

import (
	"net/http"

	"github.com/fredoliveira-ca/book-service/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.FetchData)
	http.HandleFunc("/new-book", controller.Create)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit-book", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
