package router

import (
	"net/http"

	"example.com/connectfour/internal/health"
	"example.com/connectfour/internal/ws"
)

// RegisterRoutes registers all HTTP and WebSocket routes
func RegisterRoutes() {

	// Health check route
	http.HandleFunc("/health", health.Handler)

	// WebSocket game route
	http.HandleFunc("/ws", ws.HandleGameSocket)
}
