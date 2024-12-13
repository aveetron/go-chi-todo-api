package health

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := healthCheckResponse{
		Message: "okey",
		Time:    time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
