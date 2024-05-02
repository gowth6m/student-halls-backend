package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"student-halls.com/internal/api"
)

type ShopHandler struct {
	Repo ShopRepository
}

func NewShopHandler(repo ShopRepository) *ShopHandler {
	return &ShopHandler{Repo: repo}
}

// @Summary Create a new shop
// @Description Create a new shop
// @Tags shops
// @Accept json
// @Produce json
// @Param shop body CreateShopRequest true "Shop object to be created"
// @Success 201 {object} ShopResponse "Shop created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /shop/create [post]
func (h *ShopHandler) CreateShop(c *gin.Context) {
	var reqPayload CreateShopRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoShop, err := ConvertCreateShopToShop(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdShop, err := h.Repo.CreateShop(c.Request.Context(), mongoShop)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created shop successfully", createdShop)
}

// @Summary Get all shops
// @Description Get all shops
// @Tags shops
// @Accept json
// @Produce json
// @Success 200 {object} []ShopResponse "Shops retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /shop/all [post]
func (h *ShopHandler) GetAllShops(c *gin.Context) {
	shops, err := h.Repo.GetAllShops(c.Request.Context())
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Shops retrieved successfully", shops)
}
