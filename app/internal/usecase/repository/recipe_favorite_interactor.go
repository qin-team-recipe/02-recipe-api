package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeFavoriteRepository interface {
	FindByUserID(db *gorm.DB, userID int) ([]*domain.RecipeFavorites, error)
}
