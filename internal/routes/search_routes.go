package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/services/search"
)

// GROUP: /search
func SearchRoutes(group *gin.RouterGroup) {

	// Setup the repository & handler
	searchRepo := search.NewSearchRepository()
	searchHandler := search.NewSearchHandler(searchRepo)

	searchGroup := group.Group("/search")

	// --- PUBLIC ROUTES ---
	searchGroup.GET("/", func(c *gin.Context) {
		searchHandler.Search(c)
	})
}
