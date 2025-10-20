package cmd

import (
	"context"
	"log"

	"github.com/komkem01/easy-attend-service/database/migrations"
	"github.com/uptrace/bun"
)

func modelUp(db *bun.DB) error {
	log.Printf("Executing model up...")

	// Execute raw queries before migration (enum types, extensions)
	for _, query := range migrations.RawBeforeQueryMigrate() {
		if _, err := db.Exec(query); err != nil {
			log.Printf("Error executing raw before query: %s", err)
			return err
		}
	}

	if len(migrations.Models()) == 0 {
		log.Printf("No models to migrate")
		return nil
	}

	// Create tables
	for _, mod := range migrations.Models() {
		if _, err := db.NewCreateTable().Model(mod).Exec(context.Background()); err != nil {
			return err
		}
	}

	// Execute raw queries after migration (indexes, constraints)
	for _, query := range migrations.RawAfterQueryMigrate() {
		if _, err := db.Exec(query); err != nil {
			log.Printf("Error executing raw after query: %s", err)
			return err
		}
	}

	// Seed initial data for genders and prefixes
	if err := migrations.SeedGendersAndPrefixes(context.Background(), db); err != nil {
		log.Printf("Error seeding genders and prefixes: %s", err)
		return err
	}

	return nil
}

func modelDown(db *bun.DB) error {
	log.Printf("Executing model down...")
	if len(migrations.Models()) == 0 {
		log.Printf("No models to migrate")
		return nil
	}
	// Reverse order for dropping tables
	models := migrations.Models()
	for i := len(models) - 1; i >= 0; i-- {
		if _, err := db.NewDropTable().Model(models[i]).IfExists().Cascade().Exec(context.Background()); err != nil {
			return err
		}
	}

	// Drop enum types after dropping tables
	enumTypes := []string{
		"DROP TYPE IF EXISTS user_role CASCADE",
		"DROP TYPE IF EXISTS user_status CASCADE",
		"DROP TYPE IF EXISTS gender CASCADE",
		"DROP TYPE IF EXISTS attendance_status CASCADE",
		"DROP TYPE IF EXISTS session_status CASCADE",
		"DROP TYPE IF EXISTS session_method CASCADE",
		"DROP TYPE IF EXISTS check_in_method CASCADE",
		"DROP TYPE IF EXISTS assignment_type CASCADE",
		"DROP TYPE IF EXISTS submission_format CASCADE",
		"DROP TYPE IF EXISTS assignment_status CASCADE",
		"DROP TYPE IF EXISTS submission_status CASCADE",
		"DROP TYPE IF EXISTS message_type CASCADE",
		"DROP TYPE IF EXISTS priority_level CASCADE",
		"DROP TYPE IF EXISTS event_type CASCADE",
		"DROP TYPE IF EXISTS platform CASCADE",
		"DROP TYPE IF EXISTS file_category CASCADE",
		"DROP TYPE IF EXISTS permission_type CASCADE",
		"DROP TYPE IF EXISTS notification_type CASCADE",
		"DROP TYPE IF EXISTS delivery_status CASCADE",
		"DROP TYPE IF EXISTS delivery_channel CASCADE",
		"DROP TYPE IF EXISTS reference_type CASCADE",
		"DROP TYPE IF EXISTS metric_type CASCADE",
		"DROP TYPE IF EXISTS search_type CASCADE",
		"DROP TYPE IF EXISTS risk_level CASCADE",
		"DROP TYPE IF EXISTS data_type_setting CASCADE",
		"DROP TYPE IF EXISTS classroom_status CASCADE",
		"DROP TYPE IF EXISTS classroom_role CASCADE",
		"DROP TYPE IF EXISTS member_status CASCADE",
	}

	for _, query := range enumTypes {
		if _, err := db.Exec(query); err != nil {
			log.Printf("Warning: Error dropping enum type: %s", err)
			// Continue anyway as enum types might not exist
		}
	}

	return nil
}
