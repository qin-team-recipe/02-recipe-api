package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingMemoRepository struct{}

func (sr *ShoppingMemoRepository) FirstByID(db *gorm.DB, id int) (*domain.ShoppingMemos, error) {
	s := &domain.ShoppingMemos{}
	if err := db.First(s, id).Error; err != nil {
		return &domain.ShoppingMemos{}, fmt.Errorf("not found: %w", err)
	}
	return s, nil
}

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

func (sr *ShoppingMemoRepository) Save(db *gorm.DB, s *domain.ShoppingMemos) (*domain.ShoppingMemos, error) {
	if err := db.Save(s).Error; err != nil {
		return &domain.ShoppingMemos{}, fmt.Errorf("failed shoppingMemo save: %w", err)
	}
	return s, nil
}

func (sr *ShoppingMemoRepository) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.ShoppingMemos{}, id).Error
}
