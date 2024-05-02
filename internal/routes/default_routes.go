package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student-halls.com/internal/api"
	"student-halls.com/internal/config"
)

// For /v0/
func DefaultRoutes(group *gin.RouterGroup) {
	group.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Student Halls", gin.H{
			"message": "Welcome to the Student Halls REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
		})
	})
}

// For /
func NoVersionRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Student Halls", gin.H{
			"message": "Welcome to the Student Halls REST API v0",
			"version": config.AppConfig().App.AppVersion,
			"author":  "Gowthaman Ravindrathas",
			"api":     "v0",
		})
	})
}
