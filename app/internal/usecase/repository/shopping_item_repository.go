package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingItemRepository interface {
	FirstByID(db *gorm.DB, id int) (*domain.ShoppingItems, error)
	FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.ShoppingItems, error)
	Create(db *gorm.DB, s *domain.ShoppingItems) (*domain.ShoppingItems, error)
	Save(db *gorm.DB, s *domain.ShoppingItems) (*domain.ShoppingItems, error)
	Delete(db *gorm.DB, id int) error
}
