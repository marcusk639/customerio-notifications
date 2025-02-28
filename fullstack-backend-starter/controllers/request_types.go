package controllers

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateNotificationPreferenceRequest struct {
	Type    string `json:"type" binding:"required,oneof=EMAIL SMS PUSH"`
	Enabled *bool  `json:"enabled" binding:"required"`
}
