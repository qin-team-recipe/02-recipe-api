package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRecipeRepository interface {
	FirstByRecipeID(db *gorm.DB, recipeID int) (*domain.UserRecipes, error)
	Create(db *gorm.DB, userRecipe *domain.UserRecipes) (*domain.UserRecipes, error)
}
