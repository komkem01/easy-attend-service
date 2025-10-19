package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"easy-attend-service/internal/models"
)

// Migrate runs the database migrations
func Migrate() error {
	db := GetDB()

	log.Println("Running database migrations...")

	// First, run SQL migrations for custom types
	if err := runSQLMigrations(); err != nil {
		log.Printf("SQL migration warning (may be normal if types exist): %v", err)
	}

	// Then run GORM auto-migrations
	err := db.AutoMigrate(
		// Core tables
		&models.School{},
		&models.User{},
		&models.Classroom{},
		&models.ClassroomStudent{},

		// Attendance
		&models.AttendanceSession{},
		&models.AttendanceRecord{},
		&models.AttendanceAnalytics{},

		// Assignments
		&models.Assignment{},
		&models.AssignmentSubmission{},

		// Files
		&models.FileUpload{},
		&models.FilePermission{},

		// Communication
		&models.Message{},
		&models.Notification{},

		// Calendar
		&models.AcademicCalendar{},

		// System
		&models.SystemSetting{},
		&models.SessionToken{},
	)

	if err != nil {
		return fmt.Errorf("GORM migration failed: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// runSQLMigrations runs SQL files from migrations directory
func runSQLMigrations() error {
	db := GetDB()

	// Read and execute enum creation SQL
	sqlFile := filepath.Join("../../migrations", "001_create_enums.sql")
	sqlBytes, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		// Try alternative path
		sqlFile = filepath.Join("migrations", "001_create_enums.sql")
		sqlBytes, err = ioutil.ReadFile(sqlFile)
		if err != nil {
			return fmt.Errorf("failed to read SQL migration file: %w", err)
		}
	}

	if err := db.Exec(string(sqlBytes)).Error; err != nil {
		return fmt.Errorf("failed to execute SQL migration: %w", err)
	}

	log.Println("SQL migrations completed")
	return nil
}
