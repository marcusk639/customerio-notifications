package models

type Notification struct {
	ID          string `json:"id" db:"id"`
	UserID      string `json:"userId" db:"userId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Type        string `json:"type" db:"type"`
	CreatedAt   string `json:"createdAt" db:"createdAt"`
}
