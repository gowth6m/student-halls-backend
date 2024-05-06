package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/config"
)

func SetupRoutes(router *gin.Engine) {
	// No version routes
	SwaggerRoutes(router)
	NoVersionRoutes(router)

	// Version controlled routes
	versionControlled := router.Group("/" + config.AppConfig().App.ApiVersion)
	DefaultRoutes(versionControlled)
	UserRoutes(versionControlled)
	UniversityRoutes(versionControlled)
}
