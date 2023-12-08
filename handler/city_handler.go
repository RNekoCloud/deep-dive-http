package handler

import (
	"encoding/json"
	"net/http"
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
