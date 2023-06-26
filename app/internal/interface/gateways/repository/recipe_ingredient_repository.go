package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeIngredientRepository struct{}

func (rr *RecipeIngredientRepository) FirstByID(db *gorm.DB, id int) (*domain.RecipeIngredients, error) {
	r := &domain.RecipeIngredients{}
	if err := db.First(r, id).Error; err != nil {
		return &domain.RecipeIngredients{}, fmt.Errorf("not found: %w", err)
	}
	return r, nil
}

func (rr *RecipeIngredientRepository) Create(db *gorm.DB, r *domain.RecipeIngredients) (*domain.RecipeIngredients, error) {
	if err := db.Create(r).Error; err != nil {
		return &domain.RecipeIngredients{}, fmt.Errorf("failed recipeIngredient create: %w", err)
	}
	return r, nil
}
