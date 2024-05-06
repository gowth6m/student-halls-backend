package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"student-halls.com/internal/config"
	"student-halls.com/internal/db"
	"student-halls.com/internal/routes"
)

// @title Student Halls API
// @version 1
// @description This is the REST API for Student Halls.
// @host api.student-halls.com
// @BasePath /v0
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	defer db.DisconnectFromMongoDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	routes.SetupRoutes(router)
	router.Run(config.AppConfig().App.Host + ":" + config.AppConfig().App.Port)
}
