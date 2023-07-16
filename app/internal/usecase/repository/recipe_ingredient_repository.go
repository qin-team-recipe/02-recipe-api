package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	FirstByID(db *gorm.DB, id int) (*domain.RecipeIngredients, error)
	FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.RecipeIngredients, error)
	Create(db *gorm.DB, r *domain.RecipeIngredients) (*domain.RecipeIngredients, error)
}
