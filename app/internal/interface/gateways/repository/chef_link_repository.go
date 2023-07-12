package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefLinkRepository struct{}

func (cr *ChefLinkRepository) FindByChefID(db *gorm.DB, chefID int) ([]*domain.ChefLinks, error) {
	links := []*domain.ChefLinks{}
	db.Where("chef_id = ?", chefID).Find(&links)
	if len(links) <= 0 {
		return []*domain.ChefLinks{}, fmt.Errorf("not found: %w", errors.New("chefLinks is not found"))
	}
	return links, nil
}

func (cr *ChefLinkRepository) Create(db *gorm.DB, link *domain.ChefLinks) (*domain.ChefLinks, error) {
	if err := db.Create(link).Error; err != nil {
		return &domain.ChefLinks{}, fmt.Errorf("failed cheflink create: %w", err)
	}
	return link, nil
}
