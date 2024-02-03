package routes

import (
	"github.com/go-chi/chi"
)

const (
	healthCheck    = "/health"
	getPlayerStats = "/player/{playerTag}"
	getClubStats   = "/club/{clubTag}"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get(healthCheck, healthHandler)
	r.Get(getPlayerStats, playerHandler)
	// r.Get(getClubStats, clubHandler)

	return r
}
