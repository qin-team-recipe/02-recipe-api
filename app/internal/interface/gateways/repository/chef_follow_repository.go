package repository

import (
	"errors"
	"fmt"
	"time"

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

func (cr *ChefFollowRepository) FindBybyNumberOfFollowSubscriptions(db *gorm.DB) (map[int]int64, error) {

	currentTime := time.Now().Unix()
	beforeCurrentTime := time.Now().AddDate(0, 0, -30).Unix()

	type Result struct {
		ChefID int
		Count  int64
	}

	results := []Result{}

	if err := db.Table("chef_follows").Select("chef_id, count(chef_id) as count").Where("? < created_at and created_at < ?", beforeCurrentTime, currentTime).Group("chef_id").Limit(5).Find(&results).Error; err != nil {

	}

	counts := map[int]int64{}

	for _, result := range results {
		counts[result.ChefID] = result.Count
	}

	if len(counts) <= 0 {
		return map[int]int64{}, errors.New("見つかりません")
	}

	return counts, nil
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

func (rr *ChefFollowRepository) CountByChefID(db *gorm.DB, chefID int) int {
	var count int64
	db.Model(&domain.ChefFollows{}).Where("chef_id = ?", chefID).Count(&count)
	return int(count)
}

func (cr *ChefFollowRepository) Delete(db *gorm.DB, follow *domain.ChefFollows) error {
	if err := db.Delete(follow).Error; err != nil {
		return fmt.Errorf("failed chef follow: %w", err)
	}
	return nil
}
