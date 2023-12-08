package main

import (
	"fmt"
	"net/http"
)

type City struct {
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
}

var store interface{}

func getCity(w http.ResponseWriter, req *http.Request) {

}

func main() {
	http.HandleFunc("/", getCity)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Failed to listening server:", err)
		return
	}
}
