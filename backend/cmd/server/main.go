package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/connectfour/internal/router"
	"example.com/connectfour/internal/db"
	"example.com/connectfour/internal/httpapi"
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

	http.ListenAndServe(":"+port, nil)
}
