package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	Find(db *gorm.DB) ([]*domain.Recipes, error)
	FindByUserID(db *gorm.DB, userID int) ([]*domain.Recipes, error)
}
