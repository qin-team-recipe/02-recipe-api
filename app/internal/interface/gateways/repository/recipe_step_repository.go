package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeStepRepository struct{}

func (rr *RecipeStepRepository) FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.RecipeSteps, error) {
	r := []*domain.RecipeSteps{}
	db.Where("recipe_id = ?", recipeID).Find(&r)
	if len(r) < 0 {
		return []*domain.RecipeSteps{}, errors.New("recipeStep is not found")
	}
	return r, nil
}

func (rr *RecipeStepRepository) Create(db *gorm.DB, r *domain.RecipeSteps) (*domain.RecipeSteps, error) {
	if err := db.Create(r).Error; err != nil {
		return &domain.RecipeSteps{}, fmt.Errorf("failed recipeStep create: %w", err)
	}
	return r, nil
}
