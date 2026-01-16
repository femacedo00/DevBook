package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorAPI represents the error response from API
type ErrorAPI struct {
	Error string `json:"error"`
}

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if error := json.NewEncoder(w).Encode(data); error != nil {
			log.Fatal(error)
		}
	}
}

// handleErrorStatusCode treats all request with status bigger or equal to 400
func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var errorAPI ErrorAPI
	json.NewDecoder(r.Body).Decode(&errorAPI)
	JSON(w, r.StatusCode, errorAPI)
}
