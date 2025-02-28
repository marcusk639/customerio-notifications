package services

import "github.com/customerio/fullstack-backend-homework/repositories"

type NotificationPreferenceService struct {
	repository *repositories.NotificationPreferenceRepository
}

func NewNotificationPreferenceService(repository *repositories.NotificationPreferenceRepository) *NotificationPreferenceService {
	return &NotificationPreferenceService{repository: repository}
}

func (s *NotificationPreferenceService) UpdateNotificationPreference(userID, notificationType string, enabled bool) error {
	return s.repository.UpdateNotificationPreference(userID, notificationType, enabled)
}
