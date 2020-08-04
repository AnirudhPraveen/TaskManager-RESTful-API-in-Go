func StartUp() {
	// Initialize AppConfig variable
	initConfig()
	// Initialize private/public keys for JWT authentication
	intiKeys()
	// Start a MOngoDB session
	createDbSession()
	// Add indexes into MongoDB
	addIndexes()
}