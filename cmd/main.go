package main

import (
	"encoding/json"
	"log"
	"os"
)

type MockData struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func main() {
	newData := MockData{
		Name: "John Doe",
		Age:  10,
	}

	b, _ := json.MarshalIndent(&newData, "", "	")

	err := os.WriteFile("data.json", b, 0644)

	if err != nil {
		log.Print("Failed to write file:", err)
	}
}
