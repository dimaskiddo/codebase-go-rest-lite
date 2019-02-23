package service

import (
	"log"
)

// Initialize Function in Utils
func Initialize() {
	// Initialize Configuration
	log.Println("Initialize - Config")
	initConfig()

	// Initialize Router
	log.Println("Initialize - Router")
	initRouter()
}
