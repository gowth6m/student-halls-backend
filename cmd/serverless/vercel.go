package handler

import (
	"log"
	"net/http"
	"student-halls.com/internal/config"
	"student-halls.com/internal/db"
	"student-halls.com/internal/routes"
)

// @title Student Halls API
// @version 0
// @description This is the REST API for Student Halls.
// @host api.student-halls.com
// @BasePath /v0
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Handler(w http.ResponseWriter, r *http.Request) {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	defer db.DisconnectFromMongoDB()

	router := routes.SetupRoutes()
	router.ServeHTTP(w, r)
}
