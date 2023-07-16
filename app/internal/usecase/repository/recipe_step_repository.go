package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeStepRepository interface {
	FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.RecipeSteps, error)
	Create(db *gorm.DB, r *domain.RecipeSteps) (*domain.RecipeSteps, error)
}
