package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"student-halls.com/internal/api"
)

type SearchHandler struct {
	Repo SearchRepository
}

func NewSearchHandler(repo SearchRepository) *SearchHandler {
	return &SearchHandler{Repo: repo}
}

// @Summary Search for universities and halls
// @Description Search for universities and halls
// @Tags search
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {object} SearchResult "Search results retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		api.Error(c, http.StatusBadRequest, "query parameter is required")
		return
	}

	criteria := SearchCriteria{Query: query}

	result, err := h.Repo.Search(c.Request.Context(), criteria)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Search results retrieved successfully", result)
}
