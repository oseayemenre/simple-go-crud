package handlers

import (
	"net/http"

	"github.com/oseayemenre/go_crud_scratch/internal/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.WriteToJSON(w, 200, struct{Message string `json:"message"`}{Message: "Server is healthy"})
}