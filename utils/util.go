package utils

import (
	"encoding/json"
	"net/http"
)

// Message - custom return message
func Message(status interface{}, message string) map[string]interface{} {
	return map[string]interface{}{"Status": status, "message": message}
}

// Respond - Add Header application/json and Encode data
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// CheckHTTPMethod func
func CheckHTTPMethod(method string, r *http.Request) bool {
	if method != "GET" && method != "POST" && method != "PUT" && method != "DELETE" {
		return false
	}
	if method != r.Method {
		return false
	}
	return true
}
