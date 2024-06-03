// brawl-stars-stats-service/internal/routes/routes.go
package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timolung/brawl-stars-stats-service/internal/config"
)

// NewRouter creates a new Router instance with the given configuration
func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/battle-log-stats/{playerTag}", handleBattleLogStats).Methods("GET")
	r.HandleFunc("/health-check", handleHealthCheck).Methods("GET")
	return r
}

func handleBattleLogStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerTag := vars["playerTag"]

	// Construct the URL
	url := fmt.Sprintf(config.BattleLogEndpoint, playerTag)

	// Make a request to the external endpoint
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Unexpected response from external API: %v", resp.Status), http.StatusInternalServerError)
		return
	}

	// Copy the response from the external API to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, fmt.Sprintf("Error copying response: %v", err), http.StatusInternalServerError)
		return
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Health Check OK")
}
