package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRecipeRepository struct{}

func (cr *ChefRecipeRepository) Create(db *gorm.DB, chefRecipe *domain.ChefRecipes) (*domain.ChefRecipes, error) {
	if err := db.Create(chefRecipe).Error; err != nil {
		return &domain.ChefRecipes{}, fmt.Errorf("failed chefRecipe create: %w", err)
	}
	return chefRecipe, nil
}
