package routes

import "github.com/gin-gonic/gin"
import "student-halls.com/internal/services/university"

// GROUP: /university
func UniversityRoutes(group *gin.RouterGroup) {

	// Setup the repository & handler
	universityRepo := university.NewUniversityRepository()
	universityHandler := university.NewUniversityHandler(universityRepo)

	universityGroup := group.Group("/university")

	// --- PUBLIC ROUTES ---
	universityGroup.POST("/create", func(c *gin.Context) {
		universityHandler.CreateUniversity(c)
	})

	universityGroup.GET("/all", func(c *gin.Context) {
		universityHandler.GetAllUniversities(c)
	})

	universityGroup.GET("/:id", func(c *gin.Context) {
		universityHandler.GetUniversityByID(c)
	})

	universityGroup.GET("/:id/halls", func(c *gin.Context) {
		universityHandler.GetHallsByUniversityID(c)
	})
}
