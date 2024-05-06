package serverless

import (
	"github.com/gin-contrib/cors"
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

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://student-halls.com", "https://www.student-halls.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	return router, func() {
		db.DisconnectFromMongoDB()
	}
}
