package controllers

import (
	"github.com/HardwareAndro/go-kanban-api/internal/api/services"
	models "github.com/HardwareAndro/go-kanban-api/internal/models"
	"github.com/HardwareAndro/go-kanban-api/internal/shared/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_INPUT})
		return
	}
	createdUser, err := uc.userService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_REGISTER_USER})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

// LoginUser handles user login
func (uc *UserController) LoginUser(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ERR_INVALID_INPUT})
		return
	}
	user, err := uc.userService.LoginUser(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": constants.ERR_INVALID_CREDENTIALS})
		return
	}
	c.JSON(http.StatusOK, user)
}

// LogoutUser handles user logout
func (uc *UserController) LogoutUser(c *gin.Context) {
	if err := uc.userService.LogoutUser(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ERR_LOGOUT_FAILED})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
