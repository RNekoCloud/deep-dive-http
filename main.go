package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type City struct {
	Name     string `json:"name,omitempty"`
	Province string `json:"province,omitempty"`
}

var store interface{}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	res := map[string]interface{}{
		"message": "City API Version 1.0.0",
		"status":  "success",
	}

	json, _ := json.Marshal(&res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(json)

}

func main() {
	// Loggin server listener
	fmt.Println("Server is running on http://localhost:3000/")

	// Route
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Failed to listening server:", err)
		return
	}
}
