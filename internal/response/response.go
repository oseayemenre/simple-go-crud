package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteToJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")	
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Fatalf("error: %v", err)
	}
}