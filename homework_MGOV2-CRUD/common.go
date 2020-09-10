package main

import (
	"encoding/json"
	"net/http"
)

func respondJSON(res http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	res.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
