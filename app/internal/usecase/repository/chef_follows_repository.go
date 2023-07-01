package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefFollowRepository interface {
	FindByUserID(db *gorm.DB, userID int) ([]*domain.ChefFollows, error)
	FirstByUserIDAndChefID(db *gorm.DB, userID, chefID int) (*domain.ChefFollows, error)
	Create(db *gorm.DB, follow *domain.ChefFollows) (*domain.ChefFollows, error)
	Delete(db *gorm.DB, follow *domain.ChefFollows) error
}
