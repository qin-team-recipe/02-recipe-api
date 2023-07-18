package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRecipeRepository interface {
	FirstByRecipeID(db *gorm.DB, recipeID int) (*domain.ChefRecipes, error)
	FindInByChefIDs(db *gorm.DB, ids []int) ([]*domain.ChefRecipes, error)
	Create(db *gorm.DB, chefRecipe *domain.ChefRecipes) (*domain.ChefRecipes, error)
	CountByChefID(db *gorm.DB, chefID int) int
}
