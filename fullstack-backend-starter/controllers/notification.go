// api/handlers.go
package controllers

import (
	"net/http"

	"github.com/customerio/fullstack-backend-homework/services"
	"github.com/gin-gonic/gin"
)

// NotificationController handles notification-related requests
type NotificationController struct {
	service *services.NotificationService
}

// NewNotificationController creates a new notification controller
func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{service: service}
}

// GetNotifications handles GET /notifications requests
func (h *NotificationController) GetNotifications(c *gin.Context) {
	userID := c.GetString("user_id")

	response, err := h.service.GetUserNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
