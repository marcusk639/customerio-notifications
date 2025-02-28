package repositories

import (
	"github.com/jmoiron/sqlx"
)

type NotificationPreferenceRepository struct {
	Repository
}

func NewNotificationPreferenceRepository(db *sqlx.DB) *NotificationPreferenceRepository {
	return &NotificationPreferenceRepository{
		Repository: Repository{db: db},
	}
}

// UpdateNotificationPreference updates a user's notification preference
func (r *NotificationPreferenceRepository) UpdateNotificationPreference(userID, notificationType string, enabled bool) error {
	_, err := r.db.Exec(`
		INSERT INTO notification_preferences (userId, type, enabled, updatedAt)
		VALUES (?, ?, ?, strftime('%Y-%m-%d %H:%M:%S', 'now'))
		ON CONFLICT(userId, type) DO UPDATE SET 
			enabled = ?,
			updatedAt = strftime('%Y-%m-%d %H:%M:%S', 'now')
	`, userID, notificationType, enabled, enabled)
	return err
}
