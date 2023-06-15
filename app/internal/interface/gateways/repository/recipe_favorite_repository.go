package repository

import (
	"errors"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeFavoriteRepository struct{}

func (rr *RecipeFavoriteRepository) FindByUserID(db *gorm.DB, userID int) ([]*domain.RecipeFavorites, error) {
	recipeFavorites := []*domain.RecipeFavorites{}
	db.Where("user_id = ?", userID).Find(&recipeFavorites)
	if len(recipeFavorites) <= 0 {
		return []*domain.RecipeFavorites{}, errors.New("recipeFavorites is not found")
	}
	return recipeFavorites, nil
}
