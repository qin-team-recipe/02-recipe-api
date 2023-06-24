package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeLinkRepository interface {
	Create(db *gorm.DB, r *domain.RecipeLinks) (*domain.RecipeLinks, error)
}
