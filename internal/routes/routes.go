// brawl-stars-stats-service/internal/routes/routes.go
package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timolung/brawl-stars-stats-service/internal/services"
)

// NewRouter creates a new Router instance with the given configuration
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/battle-log-stats/{playerTag}", getBattleLogStats)
	r.HandleFunc("/health", getHealthCheck).Methods("GET")
	return r
}

func getBattleLogStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerTag := vars["playerTag"]
	playerService := services.NewPlayerService(playerTag)
	data, err := playerService.GetData()
	if err != nil {
		http.Error(w, "Failed to fetch player data", http.StatusInternalServerError)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the data as JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Health Check OK")
}
