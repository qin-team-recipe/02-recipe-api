package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRecipeRepository interface {
	FirstByRecipeID(db *gorm.DB, recipeID int) (*domain.ChefRecipes, error)
	Create(db *gorm.DB, chefRecipe *domain.ChefRecipes) (*domain.ChefRecipes, error)
}
