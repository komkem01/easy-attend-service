package migrations

import (
	"context"
	"fmt"

	"github.com/komkem01/easy-attend-service/model"
	"github.com/uptrace/bun"
)

// SeedGendersAndPrefixes seeds initial data for genders and prefixes tables
func SeedGendersAndPrefixes(ctx context.Context, db *bun.DB) error {
	// Seed Genders
	genders := []model.Genders{
		{
			ID:           1,
			Code:         "M",
			NameTH:       "ชาย",
			NameEN:       "Male",
			Abbreviation: "M",
			IsActive:     true,
			SortOrder:    1,
		},
		{
			ID:           2,
			Code:         "F",
			NameTH:       "หญิง",
			NameEN:       "Female",
			Abbreviation: "F",
			IsActive:     true,
			SortOrder:    2,
		},
		{
			ID:           3,
			Code:         "O",
			NameTH:       "อื่นๆ",
			NameEN:       "Other",
			Abbreviation: "O",
			IsActive:     true,
			SortOrder:    3,
		},
	}

	for _, gender := range genders {
		_, err := db.NewInsert().
			Model(&gender).
			On("CONFLICT (code) DO NOTHING").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to seed gender %s: %w", gender.Code, err)
		}
	}

	// Seed Prefixes
	prefixes := []model.Prefixes{
		{
			ID:           1,
			Code:         "MR",
			NameTH:       "นาย",
			NameEN:       "Mr.",
			Abbreviation: "นาย",
			GenderCode:   stringPtr("M"),
			IsActive:     true,
			SortOrder:    1,
		},
		{
			ID:           2,
			Code:         "MRS",
			NameTH:       "นาง",
			NameEN:       "Mrs.",
			Abbreviation: "นาง",
			GenderCode:   stringPtr("F"),
			IsActive:     true,
			SortOrder:    2,
		},
		{
			ID:           3,
			Code:         "MISS",
			NameTH:       "นางสาว",
			NameEN:       "Miss",
			Abbreviation: "น.ส.",
			GenderCode:   stringPtr("F"),
			IsActive:     true,
			SortOrder:    3,
		},
		{
			ID:           4,
			Code:         "DR",
			NameTH:       "ดร.",
			NameEN:       "Dr.",
			Abbreviation: "ดร.",
			GenderCode:   nil, // สำหรับทุกเพศ
			IsActive:     true,
			SortOrder:    4,
		},
		{
			ID:           5,
			Code:         "PROF",
			NameTH:       "ศ.",
			NameEN:       "Prof.",
			Abbreviation: "ศ.",
			GenderCode:   nil, // สำหรับทุกเพศ
			IsActive:     true,
			SortOrder:    5,
		},
		{
			ID:           6,
			Code:         "ASSOC_PROF",
			NameTH:       "รศ.",
			NameEN:       "Assoc. Prof.",
			Abbreviation: "รศ.",
			GenderCode:   nil, // สำหรับทุกเพศ
			IsActive:     true,
			SortOrder:    6,
		},
		{
			ID:           7,
			Code:         "ASST_PROF",
			NameTH:       "ผศ.",
			NameEN:       "Asst. Prof.",
			Abbreviation: "ผศ.",
			GenderCode:   nil, // สำหรับทุกเพศ
			IsActive:     true,
			SortOrder:    7,
		},
	}

	for _, prefix := range prefixes {
		_, err := db.NewInsert().
			Model(&prefix).
			On("CONFLICT (code) DO NOTHING").
			Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed to seed prefix %s: %w", prefix.Code, err)
		}
	}

	fmt.Println("✅ Genders and Prefixes seeded successfully")
	return nil
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}
