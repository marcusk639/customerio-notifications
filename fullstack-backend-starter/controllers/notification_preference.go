package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/customerio/fullstack-backend-homework/services"
)

// NotificationPreferenceController handles notification preference-related requests
type NotificationPreferenceController struct {
	service *services.NotificationPreferenceService
}

// NewNotificationPreferenceController creates a new notification preference controller
func NewNotificationPreferenceController(service *services.NotificationPreferenceService) *NotificationPreferenceController {
	return &NotificationPreferenceController{service: service}
}

// UpdatePreferences handles POST /preferences requests
func (h *NotificationPreferenceController) UpdatePreferences(c *gin.Context) {
	userID := c.GetString("user_id")

	var request UpdateNotificationPreferenceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateNotificationPreference(userID, request.Type, *request.Enabled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
