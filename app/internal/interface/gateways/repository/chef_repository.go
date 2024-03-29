package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRepository struct{}

func (cr *ChefRepository) Find(db *gorm.DB) ([]*domain.Chefs, error) {
	chefs := []*domain.Chefs{}
	if err := db.Find(&chefs).Error; err != nil {
		return []*domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
	}
	return chefs, nil
}

func (cr *ChefRepository) FindByQuery(db *gorm.DB, q string, cursor, limit int) ([]*domain.Chefs, error) {
	chefs := []*domain.Chefs{}
	query := db.Where("? < id", cursor).Limit(limit).Order("created_at desc")

	if q != "" {
		q = "%" + q + "%"

		query = query.Where("display_name like ? or description like ?", q, q)
	}

	if err := query.Find(&chefs).Error; err != nil {
		return []*domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
	}
	return chefs, nil
}

func (cr *ChefRepository) FindInChefIDs(db *gorm.DB, ids []int) ([]*domain.Chefs, error) {
	chefs := []*domain.Chefs{}
	db.Where("id in ?", ids).Find(&chefs)
	if len(chefs) <= 0 {
		return []*domain.Chefs{}, fmt.Errorf("not found: %w", errors.New("chefs is not found"))
	}
	return chefs, nil
}

func (cr *ChefRepository) FirstByID(db *gorm.DB, id int) (*domain.Chefs, error) {
	chef := &domain.Chefs{}
	if err := db.First(chef, id).Error; err != nil {
		return &domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
	}
	return chef, nil
}

func (cr *ChefRepository) FirstByScreenName(db *gorm.DB, screenName string) (*domain.Chefs, error) {
	chef := &domain.Chefs{}
	if err := db.Where("screen_name = ?", screenName).First(chef).Error; err != nil {
		return &domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
	}
	return chef, nil
}

func (cr *ChefRepository) ExistsByScreenName(db *gorm.DB, screenName string) (bool, error) {
	var count int64 = 0
	chef := &domain.Chefs{}
	if err := db.Where("screen_name = ?", screenName).Find(chef).Count(&count).Error; err != nil {
		return true, fmt.Errorf("query error: %w", err)
	}
	if int(count) > 1 {
		return true, nil
	}
	return false, nil
}

func (cr *ChefRepository) Create(db *gorm.DB, chef *domain.Chefs) (*domain.Chefs, error) {
	if err := db.Create(chef).Error; err != nil {
		return &domain.Chefs{}, fmt.Errorf("failed chef create: %w", err)
	}
	return chef, nil
}
