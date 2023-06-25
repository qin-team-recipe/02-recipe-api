package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingMemoRepository struct{}

func (sr *ShoppingMemoRepository) FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.ShoppingMemos, error) {
	s := []*domain.ShoppingMemos{}
	db.Joins("join recipe_ingredients on recipe_ingredients.id = shopping_memos.recipe_ingredient_id").Where("recipe_ingredients.recipe_id = ?", recipeID).Find(&s)
	if len(s) < 0 {
		return []*domain.ShoppingMemos{}, errors.New("shoppingMemos is not found")
	}
	return s, nil
}

func (sr *ShoppingMemoRepository) Create(db *gorm.DB, s *domain.ShoppingMemos) (*domain.ShoppingMemos, error) {
	if err := db.Create(s).Error; err != nil {
		return &domain.ShoppingMemos{}, fmt.Errorf("failed shoppingMemo create: %w", err)
	}
	return s, nil
}
