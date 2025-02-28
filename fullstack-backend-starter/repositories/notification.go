// repository/repository.go
package repositories

import (
	"github.com/customerio/fullstack-backend-homework/models"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository struct {
	Repository
}

func NewNotificationRepository(db *sqlx.DB) *NotificationRepository {
	return &NotificationRepository{
		Repository: Repository{db: db},
	}
}

// GetUserNotifications retrieves notifications for a user filtered by preferences
func (r *NotificationRepository) GetUserNotifications(userID string) ([]models.Notification, error) {
	notifications := []models.Notification{}
	query := `
		SELECT n.id, n.userId, n.title, n.description, n.type, n.createdAt
		FROM notifications n
		INNER JOIN notification_preferences p 
			ON n.userId = p.userId 
			AND n.type = p.type
			AND p.enabled = 1
		WHERE n.userId = ?
		ORDER BY n.createdAt DESC
	`
	err := r.db.Select(&notifications, query, userID)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}

// NotificationTypeExists checks if a notification type exists
func (r *NotificationRepository) NotificationTypeExists(notificationType string) (bool, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM notifications WHERE type = ?", notificationType)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
