package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ShoppingMemoRepository struct{}

func (sr *ShoppingMemoRepository) Create(db *gorm.DB, s *domain.ShoppingMemos) (*domain.ShoppingMemos, error) {
	if err := db.Create(s).Error; err != nil {
		return &domain.ShoppingMemos{}, fmt.Errorf("failed shoppingMemo create: %w", err)
	}
	return s, nil
}
