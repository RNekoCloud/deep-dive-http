package routes

import (
	"net/http"

	"github.com/RNekoCloud/deep-dive-http/handler"
)

func CityRoutes() {
	http.HandleFunc("/", handler.Root)
	http.HandleFunc("/cities", handler.AddCity)
}
