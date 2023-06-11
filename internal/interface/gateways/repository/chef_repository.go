package repository

import (
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

func (cr *ChefRepository) FindByQuery(db *gorm.DB, q string) ([]*domain.Chefs, error) {
	chefs := []*domain.Chefs{}
	if err := db.Where("display_name = ? or description = ?", q, q).Find(chefs).Error; err != nil {
		return []*domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
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
	if err := db.Where("sceen_name = ?", screenName).First(chef).Error; err != nil {
		return &domain.Chefs{}, fmt.Errorf("chef is not found: %w", err)
	}
	return chef, nil
}
