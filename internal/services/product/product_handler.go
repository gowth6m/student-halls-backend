package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"student-halls.com/internal/api"
)

type ProductHandler struct {
	Repo ProductRepository
}

func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Product object to be created"
// @Success 201 {object} ProductResponse "Product created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /product/create [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var reqPayload CreateProductRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoProd, err := ConvertCreateProductRequestToProduct(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdProduct, err := h.Repo.CreateProduct(c.Request.Context(), mongoProd)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created product successfully", createdProduct)
}

// @Summary Create many products
// @Description Create many products
// @Tags products
// @Accept json
// @Produce json
// @Param products body []CreateProductRequest true "Products to be created"
// @Success 201 {object} []ProductResponse "Products created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /product/createMany [post]
func (h *ProductHandler) CreateManyProduct(c *gin.Context) {
	var reqPayload []CreateProductRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var mongoProds []Product

	for _, req := range reqPayload {
		mongoProd, err := ConvertCreateProductRequestToProduct(req)
		if err != nil {
			api.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		mongoProds = append(mongoProds, mongoProd)
	}

	createdProduct, err := h.Repo.CreateManyProduct(c.Request.Context(), mongoProds)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created products successfully", createdProduct)
}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} []ProductResponse "Products retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /product/all [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.Repo.GetAllProducts(c.Request.Context())
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Retrieved products successfully", products)
}
