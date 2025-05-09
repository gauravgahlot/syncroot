package handlers

import (
	"encoding/json"
	"net/http"
)

// Health handles GET /healthz requests.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]string{"status": "ok"}
	json.NewEncoder(w).Encode(res)
}
