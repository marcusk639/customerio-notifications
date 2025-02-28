package controllers

import (
	"net/http"

	"github.com/customerio/fullstack-backend-homework/services"
	"github.com/customerio/fullstack-backend-homework/utilities"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

// A login endpoint handler that takes any id and email and returns a signed JWT
func (h *AuthController) Login(c *gin.Context) {
	email := c.PostForm("email")

	user, err := h.userService.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := utilities.CreateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
