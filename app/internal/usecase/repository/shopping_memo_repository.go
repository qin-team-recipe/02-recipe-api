package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingMemoRepository interface {
	FirstByID(db *gorm.DB, id int) (*domain.ShoppingMemos, error)
	FindByRecipeID(db *gorm.DB, recipeID int) ([]*domain.ShoppingMemos, error)
	Create(db *gorm.DB, s *domain.ShoppingMemos) (*domain.ShoppingMemos, error)
	Save(db *gorm.DB, s *domain.ShoppingMemos) (*domain.ShoppingMemos, error)
	Delete(db *gorm.DB, id int) error
}
