package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/timolung/brawl-stars-stats-service/internal/service"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling health check request")
	response := HealthResponse{
		Status: "OK",
	}

	// Respond with the health status in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func playerHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the about route
	playerTag := url.PathEscape("#" + chi.URLParam(r, "playerTag"))

	playerService := service.NewPlayerService()

	playerService.GetStatsSummary(playerTag)
	playerService.GetBattleLog(playerTag)

	fmt.Printf("playerTag: %v\n", playerTag)
	w.WriteHeader(http.StatusOK)

}

func clubHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the about route
	playerTag := url.PathEscape("#" + chi.URLParam(r, "playerTag"))

	playerService := service.NewPlayerService()

	playerService.GetStatsSummary(playerTag)
	playerService.GetBattleLog(playerTag)

	fmt.Printf("playerTag: %v\n", playerTag)
	w.WriteHeader(http.StatusOK)

}
