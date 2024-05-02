package serverless

import (
	"github.com/gin-gonic/gin"
	"log"
	"student-halls.com/internal/config"
	"student-halls.com/internal/db"
	"student-halls.com/internal/routes"
)

// This function initializes the serverless application.
// It loads the configuration, connects to the MongoDB database, sets up the routes, and returns the router.
//
// This is done to expose internal functions to the serverless environment.
//
// Returns: router, a clean up function to disconnect from the MongoDB database
func Initialize() (*gin.Engine, func()) {
	err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()

	router := routes.SetupRoutes()

	return router, func() {
		db.DisconnectFromMongoDB()
	}
}
