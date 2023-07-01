package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeFavoriteRepository interface {
	FindByUserID(db *gorm.DB, userID int) ([]*domain.RecipeFavorites, error)
	FirstByUserIDAndRecipeID(db *gorm.DB, userID, recipeID int) (*domain.RecipeFavorites, error)
	Create(db *gorm.DB, favorite *domain.RecipeFavorites) (*domain.RecipeFavorites, error)
	Delete(db *gorm.DB, favorite *domain.RecipeFavorites) error
	CountByRecipeID(db *gorm.DB, recipeID int) int
}
