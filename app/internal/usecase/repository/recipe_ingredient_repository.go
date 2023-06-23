package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeIngredientRepository interface {
	Create(db *gorm.DB, r *domain.RecipeIngredients) (*domain.RecipeIngredients, error)
}
