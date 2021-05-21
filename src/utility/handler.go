package utility

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
)

// Header set header to request
func Headers(r http.Handler) http.Handler {
	headerOk := handlers.AllowedHeaders([]string{"Authorization"})
	originOk := handlers.AllowedOrigins([]string{"*"})
	methodOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headerOk, originOk, methodOk)(r)
}

// Response will return json response of http
// This func handle both error a well as success
func Response(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(payload)
}

func ReadBody(r *http.Request, data interface{}) (interface{}, error) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	return data, err
}
