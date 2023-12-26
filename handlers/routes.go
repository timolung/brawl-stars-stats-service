package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Default HomePage, Nothing To See Here.")
	fmt.Println("Endpoint Hit: homePage")
}

func playerStats(w http.ResponseWriter, r *http.Request) {
	// use http client using uri, with same bearer token, and get response object
	// call http client twice, with player info, and player battle log
	// process battle log to get stats
	// return key player info

	fmt.Fprintf(w, "Default HomePage, Nothing To See Here.")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
