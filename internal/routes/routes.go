package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/about", AboutHandler)

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the home route
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the about route
}
