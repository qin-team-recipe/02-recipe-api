package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRepository interface {
	Find(db *gorm.DB) ([]*domain.Chefs, error)
	FirstByID(db *gorm.DB, id int) (*domain.Chefs, error)
	FirstByScreenName(db *gorm.DB, screenName string) (*domain.Chefs, error)
}