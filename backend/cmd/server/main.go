package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/connectfour/internal/db"
	"example.com/connectfour/internal/httpapi"
	"example.com/connectfour/internal/router"
)

func main() {
	db.Connect()

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting HTTP server on port %s...\n", port)
	http.HandleFunc("/leaderboard", httpapi.LeaderboardHandler)

	router.RegisterRoutes()

	http.ListenAndServe("0.0.0.0:"+port, nil)
}
