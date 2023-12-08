package main

import (
	"fmt"
	"net/http"

	"github.com/RNekoCloud/deep-dive-http/routes"
)

type City struct {
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
}

type Address = string

var store interface{}

func main() {
	var addr Address = ":3000"
	// Loggin server listener
	fmt.Println("Server is running on http://localhost:3000/")

	// Route
	routes.CityRoutes()

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Println("Failed to listening server:", err)
		return
	}
}
