package university

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"student-halls.com/internal/api"
)

type UniversityHandler struct {
	Repo UniversityRepository
}

func NewUniversityHandler(repo UniversityRepository) *UniversityHandler {
	return &UniversityHandler{Repo: repo}
}

// @Summary Create a new university
// @Description Create a new university
// @Tags universities
// @Accept json
// @Produce json
// @Param university body CreateUniversityRequest true "University object to be created"
// @Success 201 {object} UniversityResponse "University created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /university/create [post]
func (h *UniversityHandler) CreateUniversity(c *gin.Context) {
	var reqPayload CreateUniversityRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoUniversity, err := ConvertCreateUniversityRequestToUniversity(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUniversity, err := h.Repo.CreateUniversity(c.Request.Context(), mongoUniversity)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created university successfully", createdUniversity)
}

// @Summary Get all universities
// @Description Get all universities
// @Tags universities
// @Accept json
// @Produce json
// @Success 200 {object} []UniversityResponse "Universities retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /university/all [get]
func (h *UniversityHandler) GetAllUniversities(c *gin.Context) {
	universities, err := h.Repo.GetAllUniversities(c.Request.Context())
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Retrieved universities successfully", universities)
}

// @Summary Get university by ID
// @Description Get university by ID
// @Tags universities
// @Accept json
// @Produce json
// @Param id path string true "University ID"
// @Success 200 {object} UniversityResponse "University retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 404 {object} map[string]interface{} "University not found"
// @Router /university/id/{id} [get]
func (h *UniversityHandler) GetUniversityByID(c *gin.Context) {
	university, err := h.Repo.GetUniversityByID(c.Request.Context(), c.Param("id"))

	if err != nil {
		api.Error(c, http.StatusNotFound, err.Error())
		return
	}

	universityRes := ConvertUniversityToUniversityResponse(*university)

	api.Success(c, http.StatusOK, "Retrieved university successfully", universityRes)
}
