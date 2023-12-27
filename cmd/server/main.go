// cmd/server/main.go
package main

import (
	"net/http"

	"github.com/timolung/brawl-stars-stats-service/internal/routes"
)

func main() {
	r := routes.SetupRoutes()

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
