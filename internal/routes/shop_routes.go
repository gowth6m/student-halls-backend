package routes

import (
	"github.com/gin-gonic/gin"
	"student-halls.com/internal/middleware"
	"student-halls.com/internal/services/shop"
)

// GROUP: /shop
func ShopRoutes(group *gin.RouterGroup) {
	// Setup the repository & handler
	shopRepo := shop.NewShopRepository()
	shopHandler := shop.NewShopHandler(shopRepo)

	shopGroup := group.Group("/shop")

	// --- PUBLIC ROUTES ---
	shopGroup.GET("/all", func(c *gin.Context) {
		shopHandler.GetAllShops(c)
	})

	// --- PROTECTED ROUTES ---
	shopGroup.Use(middleware.JWTAuthMiddleware())
	{
		shopGroup.POST("/create", func(c *gin.Context) {
			shopHandler.CreateShop(c)
		})
	}

}
