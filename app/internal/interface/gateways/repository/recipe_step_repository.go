package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeStepRepository struct{}

func (rr *RecipeStepRepository) Create(db *gorm.DB, r *domain.RecipeSteps) (*domain.RecipeSteps, error) {
	if err := db.Create(r).Error; err != nil {
		return &domain.RecipeSteps{}, fmt.Errorf("failed recipeStep create: %w", err)
	}
	return r, nil
}
