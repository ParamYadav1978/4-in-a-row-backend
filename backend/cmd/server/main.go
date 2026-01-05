package main

import (
	"fmt"
	"net/http"

	"example.com/connectfour/internal/router"
	"example.com/connectfour/internal/db"
	"example.com/connectfour/internal/httpapi"
)

func main() {
	db.Connect()
	fmt.Println("Starting HTTP server on port 8080...")
	http.HandleFunc("/leaderboard", httpapi.LeaderboardHandler)


	router.RegisterRoutes()

	http.ListenAndServe(":8080", nil)
}
