package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefRecipeRepository struct{}

func (cr *ChefRecipeRepository) FirstByRecipeID(db *gorm.DB, recipeID int) (*domain.ChefRecipes, error) {
	chefRecipe := &domain.ChefRecipes{}
	if err := db.Where("recipe_id = ?", recipeID).First(chefRecipe).Error; err != nil {
		return &domain.ChefRecipes{}, fmt.Errorf("chefRecipe is not found: %w", err)
	}
	return chefRecipe, nil
}

func (cr *ChefRecipeRepository) Create(db *gorm.DB, chefRecipe *domain.ChefRecipes) (*domain.ChefRecipes, error) {
	if err := db.Create(chefRecipe).Error; err != nil {
		return &domain.ChefRecipes{}, fmt.Errorf("failed chefRecipe create: %w", err)
	}
	return chefRecipe, nil
}

func (cr *ChefRecipeRepository) CountByChefID(db *gorm.DB, chefID int) int {
	var count int64
	db.Model(&domain.ChefRecipes{}).Where("chef_id = ?", chefID).Count(&count)
	return int(count)
}
