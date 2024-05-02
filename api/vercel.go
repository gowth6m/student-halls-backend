package handler

import (
	"net/http"
	"student-halls.com/pkg/serverless"
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
	router, cleanup := serverless.Initialize()
	defer cleanup()
	router.ServeHTTP(w, r)
}
