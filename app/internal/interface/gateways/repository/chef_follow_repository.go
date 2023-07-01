package repository

import (
	"errors"
	"fmt"

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

func (cr *ChefFollowRepository) FirstByUserIDAndChefID(db *gorm.DB, userID, chefID int) (*domain.ChefFollows, error) {
	chefFollow := &domain.ChefFollows{}
	if err := db.Where("user_id = ? and chef_id = ?", userID, chefID).First(chefFollow).Error; err != nil {
		return &domain.ChefFollows{}, errors.New("chefFollow is not found")
	}
	return chefFollow, nil
}

func (cr *ChefFollowRepository) Create(db *gorm.DB, follow *domain.ChefFollows) (*domain.ChefFollows, error) {
	if err := db.Create(follow).Error; err != nil {
		return &domain.ChefFollows{}, fmt.Errorf("failed chef follow create: %w", err)
	}
	return follow, nil
}

func (cr *ChefFollowRepository) Delete(db *gorm.DB, follow *domain.ChefFollows) error {
	if err := db.Delete(follow).Error; err != nil {
		return fmt.Errorf("failed chef follow: %w", err)
	}
	return nil
}
