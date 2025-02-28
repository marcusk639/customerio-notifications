package services

import (
	"github.com/customerio/fullstack-backend-homework/models"
	"github.com/customerio/fullstack-backend-homework/repositories"
)

type NotificationService struct {
	notificationRepository *repositories.NotificationRepository
}

func NewNotificationService(repository *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepository: repository}
}

func (s *NotificationService) GetUserNotifications(userID string) ([]models.Notification, error) {
	return s.notificationRepository.GetUserNotifications(userID)
}
