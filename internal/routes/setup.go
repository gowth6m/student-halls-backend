package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/config"
)

// Routes
func SetupRoutes() *gin.Engine {
	// Initialize the router
	router := gin.Default()

	SwaggerRoutes(router)
	NoVersionRoutes(router)

	// Version controlled routes
	versionControlled := router.Group("/" + config.AppConfig().App.ApiVersion)
	DefaultRoutes(versionControlled)
	UserRoutes(versionControlled)
	UniversityRoutes(versionControlled)

	return router
}
