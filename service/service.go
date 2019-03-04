package service

// Initialize Function in Utils
func Initialize() {
	// Initialize Logger
	initLog()
	Log("info", "service-initialize", "initialize log")

	// Initialize Configuration
	Log("info", "service-initialize", "initialize config")
	initConfig()

	// Initialize Cryptography
	Log("info", "service-initialize", "initialize crypto")
	initCrypt()

	// Initialize Router
	Log("info", "service-initialize", "initialize router")
	initRouter()
}
