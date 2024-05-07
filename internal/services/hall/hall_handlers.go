package hall

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"student-halls.com/internal/api"
)

type HallHandler struct {
	Repo HallRepository
}

func NewHallHandler(repo HallRepository) *HallHandler {
	return &HallHandler{Repo: repo}
}

// @Summary Create a new hall
// @Description Create a new hall
// @Tags halls
// @Accept json
// @Produce json
// @Param hall body CreateHallRequest true "Hall object to be created"
// @Success 201 {object} HallResponse "Hall created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /hall/create [post]
func (h *HallHandler) CreateHall(c *gin.Context) {
	var reqPayload CreateHallRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoHall, err := ConvertCreateHallRequestToHall(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdHall, err := h.Repo.CreateHall(c.Request.Context(), mongoHall)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created hall successfully", createdHall)
}

// @Summary Get all halls
// @Description Get all halls
// @Tags halls
// @Accept json
// @Produce json
// @Success 200 {object} []HallResponse "Halls retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /hall/all [get]
func (h *HallHandler) GetAllHalls(c *gin.Context) {
	halls, err := h.Repo.GetAllHalls(c.Request.Context())
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	hallResponses := []HallResponse{}
	for _, hall := range halls {
		hallResponses = append(hallResponses, ConvertHallToHallResponse(hall))
	}

	api.Success(c, http.StatusOK, "Retrieved halls successfully", hallResponses)
}

// @Summary Get hall by ID
// @Description Get hall by ID
// @Tags halls
// @Accept json
// @Produce json
// @Param id path string true "Hall ID"
// @Success 200 {object} HallResponse "Hall retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 404 {object} map[string]interface{} "Hall not found"
// @Router /hall/{id} [get]
func (h *HallHandler) GetHallByID(c *gin.Context) {
	hall, err := h.Repo.GetHallByID(c.Request.Context(), c.Param("id"))

	if err != nil {
		api.Error(c, http.StatusNotFound, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Retrieved hall successfully", ConvertHallToHallResponse(*hall))
}
