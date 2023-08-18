package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefFollowRepository interface {
	FindByUserID(db *gorm.DB, userID, cursor, limit int) ([]*domain.ChefFollows, error)
	FindBybyNumberOfFollowSubscriptions(db *gorm.DB, cursor int) (map[int]int64, error)
	FirstByUserIDAndChefID(db *gorm.DB, userID, chefID int) (*domain.ChefFollows, error)
	Create(db *gorm.DB, follow *domain.ChefFollows) (*domain.ChefFollows, error)
	CountByChefID(db *gorm.DB, chefID int) int
	Delete(db *gorm.DB, follow *domain.ChefFollows) error
}
