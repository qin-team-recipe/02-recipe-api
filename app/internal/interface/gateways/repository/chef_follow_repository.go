package repository

import (
	"errors"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type ChefFollowRepository struct{}

func (cr *ChefFollowRepository) FindByUserID(db *gorm.DB, userID int) ([]*domain.ChefFollows, error) {
	chefFollows := []*domain.ChefFollows{}
	db.Where("user_id = ?", userID).Find(&chefFollows)
	if len(chefFollows) <= 0 {
		return []*domain.ChefFollows{}, errors.New("chefFollows is not found")
	}
	return chefFollows, nil
}
