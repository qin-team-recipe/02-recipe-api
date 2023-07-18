package repository

import (
	"errors"
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

func (cr *ChefRecipeRepository) FindInByChefIDs(db *gorm.DB, ids []int) ([]*domain.ChefRecipes, error) {
	chefRecipe := []*domain.ChefRecipes{}
	db.Where("chef_id in ?", ids).Order("created_at desc").Limit(10).Find(&chefRecipe)
	if len(chefRecipe) <= 0 {
		return []*domain.ChefRecipes{}, fmt.Errorf("chefRecipe is not found: %w", errors.New("フォローしているシェフはまだレシピを作成していません"))
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
