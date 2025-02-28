package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Initialize an in-memory SQLite database
func InitDB() *sqlx.DB {
	fmt.Print("Connecting to the database... ")
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		fmt.Print("Failed to connect to DB", err)
		os.Exit(1)
	}

	createSchema(db)

	// Import seed data relative to this file's location
	ImportUsers(db, "../seed/users.csv")
	ImportNotifications(db, "../seed/notifications.csv")

	fmt.Println("Database initialization complete")

	return db
}

// Create database schema with appropriate tables and constraints
func createSchema(db *sqlx.DB) {
	// Users table - stores user information
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT NOT NULL UNIQUE
		)
	`)

	// Notifications table - stores notification data
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS notifications (
			id TEXT PRIMARY KEY,
			userId TEXT NOT NULL,
			title TEXT NOT NULL,
			description TEXT NOT NULL,
			type TEXT NOT NULL,
			createdAt TIMESTAMP NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(id)
		)
	`)

	// Create index for notifications
	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_notifications_userid_type 
		ON notifications(userId, type, createdAt)
	`)

	// Notification preferences table - stores user preferences for notification types
	db.MustExec(`
		CREATE TABLE IF NOT EXISTS notification_preferences (
			userId TEXT NOT NULL,
			type TEXT NOT NULL,
			enabled BOOLEAN DEFAULT 1,
			updatedAt TIMESTAMP NOT NULL,
			PRIMARY KEY (userId, type),
			FOREIGN KEY (userId) REFERENCES users(id)		)
	`)

	// Create index for notification preferences
	db.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_notification_prefs_userid_type 
		ON notification_preferences(userId, type, enabled)
	`)

	fmt.Println("Schema created successfully")
}

// Execute the SQL file at the given path
func ExecSQLFile(db *sqlx.DB, sqlFilePath string) {
	seedNotifications, err := os.ReadFile(sqlFilePath)
	if err != nil {
		fmt.Print("Failed to load file", sqlFilePath, err)
		os.Exit(1)
	}

	if _, err := db.Exec(string(seedNotifications)); err != nil {
		fmt.Print("Failed to execute file", sqlFilePath, err)
		os.Exit(1)
	}
}
