// db.go
package database

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Helper function to open CSV files using paths relative to this source file
func openCSVFile(filePath string) (*os.File, error) {
	// Get the directory containing the current source file
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}
	currentDir := filepath.Dir(currentFile)
	absolutePath := filepath.Join(currentDir, filePath)

	return os.Open(absolutePath)
}

// Import users from CSV file
func ImportUsers(db *sqlx.DB, filePath string) {
	// Open the CSV file
	file, err := openCSVFile(filePath)
	if err != nil {
		fmt.Printf("Failed to open users CSV file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read users CSV: %v\n", err)
		return
	}

	// Skip header row if present
	if len(records) > 0 {
		// Check if first record contains headers
		if records[0][0] == "id" || records[0][0] == "ID" {
			records = records[1:]
		}
	}

	// Begin a transaction for better performance
	tx := db.MustBegin()

	// Insert each user record
	for _, record := range records {
		if len(record) >= 2 {
			tx.MustExec(
				"INSERT INTO users (id, email) VALUES (?, ?)",
				strings.TrimSpace(record[0]),
				strings.TrimSpace(record[1]),
			)
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Failed to commit user import transaction: %v\n", err)
		return
	}

	fmt.Printf("Imported %d users\n", len(records))
}

// Import notifications from CSV file
func ImportNotifications(db *sqlx.DB, filePath string) {
	// Open the CSV file
	file, err := openCSVFile(filePath)
	if err != nil {
		fmt.Printf("Failed to open notifications CSV file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read notifications CSV: %v\n", err)
		return
	}

	// Skip header row if present
	if len(records) > 0 {
		// Check if first record contains headers
		if records[0][0] == "id" || records[0][0] == "ID" {
			records = records[1:]
		}
	}

	// Begin a transaction for better performance
	tx := db.MustBegin()

	// Insert each notification record
	importCount := 0
	for _, record := range records {
		if len(record) >= 6 {
			// Parse timestamp
			timeStr := strings.TrimSpace(record[5])
			// Remove double quotes if present
			timeStr = strings.Trim(timeStr, "\"")

			var createdAt time.Time
			var parseErr error

			// Try different time formats if necessary
			formats := []string{
				time.RFC3339,
				"2006-01-02T15:04:05Z",
				"2006-01-02 15:04:05",
			}

			for _, format := range formats {
				createdAt, parseErr = time.Parse(format, timeStr)
				if parseErr == nil {
					break
				}
			}

			// Use current time if parsing fails
			if parseErr != nil {
				createdAt = time.Now()
				// Format the current time to match the CSV format
				createdAt = createdAt.Round(time.Second)
				fmt.Printf("Warning: Could not parse time '%s', using current time: %s\n",
					timeStr, createdAt.Format("2006-01-02 15:04:05"))
			}

			tx.MustExec(
				"INSERT INTO notifications (id, userId, title, description, type, createdAt) VALUES (?, ?, ?, ?, ?, ?)",
				strings.TrimSpace(record[0]),
				strings.TrimSpace(record[1]),
				strings.TrimSpace(record[2]),
				strings.TrimSpace(record[3]),
				strings.TrimSpace(record[4]),
				createdAt,
			)
			importCount++
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Failed to commit notification import transaction: %v\n", err)
		return
	}

	fmt.Printf("Imported %d notifications\n", importCount)
}

// Initialize default preferences for all notification types
func initializePreferences(db *sqlx.DB) {
	// Get all distinct user IDs
	var userIDs []string
	err := db.Select(&userIDs, "SELECT id FROM users")
	if err != nil {
		fmt.Printf("Failed to get users: %v\n", err)
		return
	}

	// Get distinct notification types
	var types []string
	err = db.Select(&types, "SELECT DISTINCT type FROM notifications")
	if err != nil {
		fmt.Printf("Failed to get notification types: %v\n", err)
		return
	}

	// Begin a transaction for better performance
	tx := db.MustBegin()

	// Create default preferences (enabled) for each user and notification type
	preferenceCount := 0
	for _, userID := range userIDs {
		for _, notifType := range types {
			tx.MustExec(
				"INSERT OR IGNORE INTO notification_preferences (userId, type, enabled) VALUES (?, ?, ?)",
				userID, notifType, true,
			)
			preferenceCount++
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Failed to commit preferences initialization: %v\n", err)
		return
	}

	fmt.Printf("Initialized %d preferences for %d users and %d notification types\n",
		preferenceCount, len(userIDs), len(types))
}
