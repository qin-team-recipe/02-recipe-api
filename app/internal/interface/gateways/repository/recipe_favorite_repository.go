package repository

import (
	"errors"
	"fmt"
	"time"

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

func (rr *RecipeFavoriteRepository) FindBybyNumberOfFavoriteSubscriptions(db *gorm.DB) (map[int]int64, error) {
	recipeFavorites := []*domain.RecipeFavorites{}

	currentTime := time.Now().Unix()
	beforeCurrentTime := time.Now().AddDate(0, 0, -30).Unix()

	type Result struct {
		RecipeID int
		Count    int64
	}

	results := []Result{}

	if err := db.Table("recipe_favorites").Select("recipe_id, count(recipe_id) as count").Where("? < created_at and created_at < ?", beforeCurrentTime, currentTime).Group("recipe_id").Limit(5).Find(&results).Error; err != nil {
		fmt.Println("err: ", err)
	}

	counts := map[int]int64{}

	for _, result := range results {
		recipeFavorite := &domain.RecipeFavorites{
			RecipeID: result.RecipeID,
		}
		recipeFavorites = append(recipeFavorites, recipeFavorite)

		counts[result.RecipeID] = result.Count
	}

	if len(counts) <= 0 {
		return map[int]int64{}, errors.New("見つかりません")
	}
	return counts, nil
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
