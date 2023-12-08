package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	datasource "github.com/RNekoCloud/deep-dive-http/data_source"
	"github.com/RNekoCloud/deep-dive-http/helper"
)

func Root(w http.ResponseWriter, req *http.Request) {
	res := map[string]interface{}{
		"message": "City API Version 1.0.0",
		"status":  "success",
	}

	json, _ := json.Marshal(&res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(json)

}

func AddCity(w http.ResponseWriter, req *http.Request) {
	var data datasource.City

	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Failed to parse payload request:", err)
		return
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Failed to parsing JSON:", err)
		return
	}

	id := helper.UUIDGen()
	datasource.Store[id] = data

	res := map[string]interface{}{
		"status":  "success",
		"message": "Successfully add new data",
		"data":    data,
	}

	b, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}

func CitiesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		AddCity(w, req)
		return
	} else if req.Method == "POST" {
		GetAllCities(w, req)
		return
	}
}

func GetAllCities(w http.ResponseWriter, req *http.Request) {
	var allCities []datasource.City

	for _, val := range datasource.Store {
		allCities = append(allCities, val)
	}

	res := map[string]interface{}{
		"status": "success",
		"data":   allCities,
	}

	b, _ := json.Marshal(&res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(b)
}
