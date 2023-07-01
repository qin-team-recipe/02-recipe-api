package repository

import (
	"errors"
	"fmt"

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

func (rr *RecipeFavoriteRepository) FirstByUserIDAndRecipeID(db *gorm.DB, userID, recipeID int) (*domain.RecipeFavorites, error) {
	favorite := &domain.RecipeFavorites{}
	if err := db.Where("user_id = ? and recipe_id = ?", userID, recipeID).First(favorite).Error; err != nil {
		return &domain.RecipeFavorites{}, errors.New("recipe favorite is not found")
	}
	return favorite, nil
}

func (rr *RecipeFavoriteRepository) Create(db *gorm.DB, favorite *domain.RecipeFavorites) (*domain.RecipeFavorites, error) {
	if err := db.Create(favorite).Error; err != nil {
		return &domain.RecipeFavorites{}, fmt.Errorf("failed recipe favorite create: %w", err)
	}
	return favorite, nil
}

func (rr *RecipeFavoriteRepository) Delete(db *gorm.DB, favorite *domain.RecipeFavorites) error {
	if err := db.Delete(favorite).Error; err != nil {
		return fmt.Errorf("failed recipe favorite delete: %w", err)
	}
	return nil
}

func (rr *RecipeFavoriteRepository) CountByRecipeID(db *gorm.DB, recipeID int) int {
	var count int64
	db.Model(&domain.RecipeFavorites{}).Where("recipe_id = ?", recipeID).Count(&count)
	return int(count)
}
