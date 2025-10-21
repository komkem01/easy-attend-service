package auth

import (
	"context"

	"github.com/komkem01/easy-attend-service/model"
)

// GetGendersService returns all active genders
func GetGendersService(ctx context.Context) ([]model.Genders, error) {
	var genders []model.Genders

	err := db.NewSelect().
		Model(&genders).
		Where("is_active = true").
		Order("sort_order ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return genders, nil
}

func GetGendersServiceByID(ctx context.Context, id int) (*model.Genders, error) {
	var gender model.Genders

	err := db.NewSelect().
		Model(&gender).
		Where("is_active = true").
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &gender, nil
}

// GetPrefixesService returns all active prefixes
func GetPrefixesService(ctx context.Context) ([]model.Prefixes, error) {
	var prefixes []model.Prefixes

	err := db.NewSelect().
		Model(&prefixes).
		Relation("Gender").
		Where("p.is_active = true").
		Order("p.sort_order ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return prefixes, nil
}

func GetPrefixesServiceByID(ctx context.Context, id int) (*model.Prefixes, error) {
	var prefix model.Prefixes

	err := db.NewSelect().
		Model(&prefix).
		Relation("Gender").
		Where("p.is_active = true").
		Where("p.id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &prefix, nil
}

// GetPrefixesByGenderService returns prefixes filtered by gender
func GetPrefixesByGenderService(ctx context.Context, genderCode string) ([]model.Prefixes, error) {
	var prefixes []model.Prefixes

	err := db.NewSelect().
		Model(&prefixes).
		Relation("Gender").
		Where("p.is_active = true").
		Where("p.gender_code = ? OR p.gender_code IS NULL", genderCode).
		Order("p.sort_order ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return prefixes, nil
}

// FindGenderIDByName finds gender ID by name (Thai or English)
func FindGenderIDByName(ctx context.Context, genderName string) (*int, error) {
	var gender model.Genders

	err := db.NewSelect().
		Model(&gender).
		Where("is_active = true").
		Where("name_th = ? OR name_en = ? OR LOWER(name_en) = LOWER(?)",
			genderName, genderName, genderName).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &gender.ID, nil
}

// FindPrefixIDByName finds prefix ID by name (Thai or English)
func FindPrefixIDByName(ctx context.Context, prefixName string) (*int, error) {
	var prefix model.Prefixes

	err := db.NewSelect().
		Model(&prefix).
		Where("is_active = true").
		Where("name_th = ? OR name_en = ? OR abbreviation = ?",
			prefixName, prefixName, prefixName).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &prefix.ID, nil
}
