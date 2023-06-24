package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeStepRepository interface {
	Create(db *gorm.DB, r *domain.RecipeSteps) (*domain.RecipeSteps, error)
}
