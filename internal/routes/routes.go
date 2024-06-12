// brawl-stars-stats-service/internal/routes/routes.go
package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timolung/brawl-stars-stats-service/internal/services"
	"github.com/timolung/brawl-stars-stats-service/internal/utils"
)

// NewRouter creates a new Router instance with the given configuration
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/battle-log/{playerTag}", getBattleLogStats)
	r.HandleFunc("/club-members/{clubTag}", getClubStats)
	r.HandleFunc("/health", getHealthCheck).Methods("GET")
	return r
}

func getBattleLogStats(w http.ResponseWriter, r *http.Request) {
	log.Printf("Begin Battle Logs Request")
	vars := mux.Vars(r)
	playerTag := vars["playerTag"]
	playerService := services.NewPlayerService(playerTag)

	battleLogData, err := playerService.GetBattleLog()
	if err != nil {
		http.Error(w, "Failed to fetch battle log data", http.StatusInternalServerError)
		return
	}
	log.Printf("Received Battle Log Data: %v", battleLogData)

	battleLogStats, err := playerService.CalculatePlayerStats(battleLogData)
	if err != nil {
		http.Error(w, "Failed to calculate battle log data", http.StatusInternalServerError)
		return
	}
	log.Printf("Calculated Battlle Log Stats: %v", battleLogStats)

	// Add Headers
	utils.AddCORS(w)

	// Encode the data as JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(battleLogStats); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		log.Printf("Error encoding JSON response: %v", err)
		return
	}
}

func getClubStats(w http.ResponseWriter, r *http.Request) {
	log.Printf("Begin Club Stats Request")
	vars := mux.Vars(r)
	clubTag := vars["clubTag"]
	clubService := services.NewClubService(clubTag)

	clubMembersData, err := clubService.GetClubMembersList()
	if err != nil {
		http.Error(w, "Failed to fetch player data", http.StatusInternalServerError)
		return
	}
	log.Printf("Received Club Members Data: %v", clubMembersData)

	clubMemberStats, err := clubService.CalculateClubMemberStats(clubMembersData)
	if err != nil {
		http.Error(w, "Failed to calculate club member data", http.StatusInternalServerError)
		return
	}
	log.Printf("Calculated Club Member Stats: %v", clubMemberStats)

	utils.AddCORS(w)

	// Encode the data as JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(clubMemberStats); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Health Check OK")
}
