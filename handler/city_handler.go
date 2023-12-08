package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	datasource "github.com/RNekoCloud/deep-dive-http/data_source"
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

	res := map[string]interface{}{
		"message": "Successfully add new data",
		"data":    body,
	}

	b, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}
