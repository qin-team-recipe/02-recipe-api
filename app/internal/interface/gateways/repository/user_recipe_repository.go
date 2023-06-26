package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRecipeRepository struct{}

func (ur *UserRecipeRepository) Create(db *gorm.DB, userRecipe *domain.UserRecipes) (*domain.UserRecipes, error) {
	if err := db.Create(userRecipe).Error; err != nil {
		return &domain.UserRecipes{}, fmt.Errorf("failed userRecipe create: %w", err)
	}
	return userRecipe, nil
}
