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
	ProductRoutes(versionControlled)
	ShopRoutes(versionControlled)
	UserRoutes(versionControlled)

	return router
}
