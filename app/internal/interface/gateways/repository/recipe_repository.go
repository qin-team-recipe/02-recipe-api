package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeRepository struct{}

func (rr *RecipeRepository) Find(db *gorm.DB) ([]*domain.Recipes, error) {
	recipes := []*domain.Recipes{}
	db.Find(&recipes)
	if len(recipes) <= 0 {
		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
	}
	return recipes, nil
}
