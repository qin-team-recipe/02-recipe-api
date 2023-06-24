package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeLinkRepository struct{}

func (rr *RecipeLinkRepository) Create(db *gorm.DB, r *domain.RecipeLinks) (*domain.RecipeLinks, error) {
	if err := db.Create(r).Error; err != nil {
		return &domain.RecipeLinks{}, fmt.Errorf("failed recipeLink create: %w", err)
	}
	return r, nil
}
