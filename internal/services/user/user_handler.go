package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"student-halls.com/internal/api"
	"student-halls.com/internal/middleware"
)

type UserHandler struct {
	Repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User object to be created"
// @Success 201 {object} UserResponse "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/create [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var reqPayload CreateUserRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	mongoUser, err := ConvertCreateUserRequestToUser(reqPayload)
	if err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := h.Repo.CreateUser(c.Request.Context(), mongoUser)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusCreated, "Created user successfully", createdUser)
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []UserResponse "Users retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/all [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.Repo.GetAllUsers(c.Request.Context())
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Retrieved users successfully", users)
}

// @Summary Get user by username
// @Description Get user by username
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User username"
// @Success 200 {object} UserResponse "User retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := h.Repo.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		api.Error(c, http.StatusNotFound, "User not found")
		return
	}

	api.Success(c, http.StatusOK, "Retrieved user successfully", user)
}

// @Summary Get user by email
// @Description Get user by email
// @Tags users
// @Accept json
// @Produce json
// @Param email path string true "User email"
// @Success 200 {object} UserResponse "User retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.Repo.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		api.Error(c, http.StatusNotFound, "User not found")
		return
	}

	api.Success(c, http.StatusOK, "Retrieved user successfully", user)
}

// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User login credentials"
// @Success 200 {object} LoginResponse "User logged in successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or parameters"
// @Failure 401 {object} map[string]interface{} "Invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/login [post]
func (h *UserHandler) LoginUser(c *gin.Context) {
	var reqPayload LoginRequest
	if err := c.ShouldBindJSON(&reqPayload); err != nil {
		api.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.Repo.GetUserByEmail(c.Request.Context(), reqPayload.Email)
	if err != nil {
		api.Error(c, http.StatusNotFound, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPayload.Password)); err != nil {
		api.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := middleware.GenerateJWTToken(user.Email, user.Username)
	if err != nil {
		api.Error(c, http.StatusInternalServerError, "An error occurred while processing your request")
		return
	}

	userRes := UserResponse{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		University:  user.University.Hex(),
		YearOfStudy: user.YearOfStudy,
	}

	api.Success(c, http.StatusOK, "User logged in successfully", LoginResponse{
		Token: token,
		User:  userRes,
	})
}

// @Summary Get current user
// @Description Get current user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} UserResponse "User retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /user/current [get]
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	user, err := h.Repo.GetUserByEmail(c.Request.Context(), c.GetString("email"))
	if err != nil {
		api.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	api.Success(c, http.StatusOK, "Retrieved user successfully", user)
}
