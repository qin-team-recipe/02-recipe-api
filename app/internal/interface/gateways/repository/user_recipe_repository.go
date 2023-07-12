package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRecipeRepository struct{}

func (cr *UserRecipeRepository) FirstByRecipeID(db *gorm.DB, recipeID int) (*domain.UserRecipes, error) {
	userRecipe := &domain.UserRecipes{}
	if err := db.Where("recipe_id = ?", recipeID).First(userRecipe).Error; err != nil {
		return &domain.UserRecipes{}, fmt.Errorf("userRecipe is not found: %w", err)
	}
	return userRecipe, nil
}

func (ur *UserRecipeRepository) Create(db *gorm.DB, userRecipe *domain.UserRecipes) (*domain.UserRecipes, error) {
	if err := db.Create(userRecipe).Error; err != nil {
		return &domain.UserRecipes{}, fmt.Errorf("failed userRecipe create: %w", err)
	}
	return userRecipe, nil
}
