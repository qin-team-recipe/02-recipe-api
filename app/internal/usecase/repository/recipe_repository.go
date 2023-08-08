package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeRepository interface {
	FirstByWatchID(db *gorm.DB, watchID string) (*domain.Recipes, error)
	Find(db *gorm.DB) ([]*domain.Recipes, error)
	FindByQuery(db *gorm.DB, userID int, q string) ([]*domain.Recipes, error)
	// FindByChefID(db *gorm.DB, chefID int) ([]*domain.Recipes, error)
	FindByUserID(db *gorm.DB, userID int) ([]*domain.Recipes, error)
	FindInRecipeIDs(db *gorm.DB, ids []int) ([]*domain.Recipes, error)
	FirstByID(db *gorm.DB, id int) (*domain.Recipes, error)
	Create(db *gorm.DB, chef *domain.Recipes) (*domain.Recipes, error)
	Save(db *gorm.DB, recipe *domain.Recipes) (*domain.Recipes, error)
}
