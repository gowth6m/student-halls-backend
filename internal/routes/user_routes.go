package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/middleware"
	"student-halls.com/internal/services/user"
)

// GROUP: /user
func UserRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	userRepo := user.NewUserRepository()
	userHandler := user.NewUserHandler(userRepo)

	userGroup := group.Group("/user")

	// --- PUBLIC ROUTES ---
	userGroup.POST("/create", func(c *gin.Context) {
		userHandler.CreateUser(c)
	})

	userGroup.GET("/all", func(c *gin.Context) {
		userHandler.GetAllUsers(c)
	})

	userGroup.GET("/email/:email", func(c *gin.Context) {
		userHandler.GetUserByEmail(c)
	})

	userGroup.GET("/:username", func(c *gin.Context) {
		userHandler.GetUserByUsername(c)
	})

	userGroup.POST("/login", func(c *gin.Context) {
		userHandler.LoginUser(c)
	})

	// --- PROTECTED ROUTES ---
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.GET("/current", func(c *gin.Context) {
			userHandler.GetCurrentUser(c)
		})
	}

}
