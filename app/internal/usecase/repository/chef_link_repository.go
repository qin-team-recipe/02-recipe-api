package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefLinkRepository interface {
	FindByChefID(db *gorm.DB, chefID int) ([]*domain.ChefLinks, error)
	Create(db *gorm.DB, link *domain.ChefLinks) (*domain.ChefLinks, error)
}
