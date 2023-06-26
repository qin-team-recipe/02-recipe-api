package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRecipeRepository interface {
	Create(db *gorm.DB, userRecipe *domain.UserRecipes) (*domain.UserRecipes, error)
}
