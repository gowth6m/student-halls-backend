package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/services/hall"
)

// GROUP: /hall
func HallRoutes(group *gin.RouterGroup) {

	// Setup the repository & handler
	hallRepo := hall.NewHallRepository()
	hallHandler := hall.NewHallHandler(hallRepo)

	hallGroup := group.Group("/hall")

	// --- PUBLIC ROUTES ---
	hallGroup.POST("/create", func(c *gin.Context) {
		hallHandler.CreateHall(c)
	})

	hallGroup.GET("/all", func(c *gin.Context) {
		hallHandler.GetAllHalls(c)
	})

	hallGroup.GET("/id/:id", func(c *gin.Context) {
		hallHandler.GetHallByID(c)
	})
}
