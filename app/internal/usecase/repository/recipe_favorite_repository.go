package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeFavoriteRepository interface {
	FindByUserID(db *gorm.DB, userID, curosr, limit int) ([]*domain.RecipeFavorites, error)
	FindByNumberOfFavoriteSubscriptions(db *gorm.DB, cursor int) (map[int]int64, error)
	FindByChefRecipeIDsAndNumberOfFavoriteSubscriptions(db *gorm.DB, chefID int, cursor, limit int) (map[int]int64, error)
	FirstByUserIDAndRecipeID(db *gorm.DB, userID, recipeID int) (*domain.RecipeFavorites, error)
	Create(db *gorm.DB, favorite *domain.RecipeFavorites) (*domain.RecipeFavorites, error)
	Delete(db *gorm.DB, favorite *domain.RecipeFavorites) error
	CountByRecipeID(db *gorm.DB, recipeID int) int
}
