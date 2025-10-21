package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)

// AddDeletedAtColumns adds deleted_at column to tables that don't have it
func AddDeletedAtColumns(ctx context.Context, db *bun.DB) error {
	fmt.Print(" [adding deleted_at columns]")

	// Add deleted_at column to tables that don't have it yet
	tables := []string{
		"users",
		"schools",
		"genders",
		"prefixes",
		"system_settings",
		"session_tokens",
		"user_sessions",
		"user_role_permissions",
		"user_profiles",
		"notifications",
		"security_events",
		"search_logs",
		"metrics_data",
	}

	for _, table := range tables {
		// Check if column exists first
		var exists bool
		err := db.NewRaw(`
				SELECT EXISTS (
					SELECT 1 FROM information_schema.columns 
					WHERE table_name = ? AND column_name = 'deleted_at'
				)
			`, table).Scan(ctx, &exists)

		if err != nil {
			return fmt.Errorf("failed to check if deleted_at exists in %s: %w", table, err)
		}

		// Add column if it doesn't exist
		if !exists {
			_, err = db.NewRaw(fmt.Sprintf(`
					ALTER TABLE %s 
					ADD COLUMN deleted_at TIMESTAMP DEFAULT NULL
				`, table)).Exec(ctx)

			if err != nil {
				return fmt.Errorf("failed to add deleted_at to %s: %w", table, err)
			}

			fmt.Printf(" [%s.deleted_at]", table)
		}
	}

	return nil
}
