package util

import (
	"encoding/json"
	"net/http"
)

// interface{} ---> it takes any types of data: int, struct, string etc...
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	// backend theke kono information jodi JSON e pathaite chai tahole encoder
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}


func SendError(w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}