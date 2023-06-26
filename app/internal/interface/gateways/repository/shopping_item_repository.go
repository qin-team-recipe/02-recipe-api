package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingItemRepository struct{}

func (sr *ShoppingItemRepository) FirstByID(db *gorm.DB, id int) (*domain.ShoppingItems, error) {
	s := &domain.ShoppingItems{}
	if err := db.First(s, id).Error; err != nil {
		return &domain.ShoppingItems{}, fmt.Errorf("not found: %w", err)
	}
	return s, nil
}

func (sr *ShoppingItemRepository) FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.ShoppingItems, error) {
	s := []*domain.ShoppingItems{}
	db.Joins("join recipe_ingredients on recipe_ingredients.id = shopping_ShoppingItems.recipe_ingredient_id").Where("recipe_ingredients.recipe_id = ?", recipeID).Find(&s)
	if len(s) < 0 {
		return []*domain.ShoppingItems{}, errors.New("shoppingShoppingItems is not found")
	}
	return s, nil
}

func (sr *ShoppingItemRepository) Create(db *gorm.DB, s *domain.ShoppingItems) (*domain.ShoppingItems, error) {
	if err := db.Create(s).Error; err != nil {
		return &domain.ShoppingItems{}, fmt.Errorf("failed shoppingItem create: %w", err)
	}
	return s, nil
}

func (sr *ShoppingItemRepository) Save(db *gorm.DB, s *domain.ShoppingItems) (*domain.ShoppingItems, error) {
	if err := db.Save(s).Error; err != nil {
		return &domain.ShoppingItems{}, fmt.Errorf("failed shoppingItem save: %w", err)
	}
	return s, nil
}

func (sr *ShoppingItemRepository) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.ShoppingItems{}, id).Error
}
