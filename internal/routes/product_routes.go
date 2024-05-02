package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/middleware"
	"student-halls.com/internal/services/product"
)

// GROUP: /product
func ProductRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	productRepo := product.NewProductRepository()
	productHandler := product.NewProductHandler(productRepo)

	productGroup := group.Group("/product")

	// --- PUBLIC ROUTES ---
	productGroup.GET("/all", func(c *gin.Context) {
		productHandler.GetAllProducts(c)
	})

	// --- PROTECTED ROUTES ---
	productGroup.Use(middleware.JWTAuthMiddleware())
	{
		productGroup.POST("/create", func(c *gin.Context) {
			productHandler.CreateProduct(c)
		})

		productGroup.POST("/createMany", func(c *gin.Context) {
			productHandler.CreateManyProduct(c)
		})
	}

}
