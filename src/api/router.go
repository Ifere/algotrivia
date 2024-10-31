package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", HealthCheck).Methods("GET")
	router.HandleFunc("/api/version", Version).Methods("GET")
}

// HealthCheck returns a simple health check response
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// Version returns the API version
func Version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("v1"))
}