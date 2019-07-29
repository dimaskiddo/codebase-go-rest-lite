package ctl

import (
	"net/http"

	"github.com/dimaskiddo/codebase-go-rest-lite/hlp/router"
)

// GetIndex Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	router.ResponseSuccess(w, "Codebase Go REST (Lite) is running")
}

// GetHealth Function to Show Health Check Status
func GetHealth(w http.ResponseWriter, r *http.Request) {
	router.HealthCheck(w)
}
