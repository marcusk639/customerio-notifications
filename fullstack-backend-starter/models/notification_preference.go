package models

type NotificationPreference struct {
	UserID    string `json:"userId" db:"userId"`
	Type      string `json:"type" db:"type"`
	Enabled   bool   `json:"enabled" db:"enabled"`
	UpdatedAt string `json:"updatedAt" db:"updatedAt"`
}
