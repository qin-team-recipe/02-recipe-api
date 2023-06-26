package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserShoppingItemRepository struct{}

func (ur *UserShoppingItemRepository) FindByUserID(db *gorm.DB, userID int) ([]*domain.UserShoppingItems, error) {
	items := []*domain.UserShoppingItems{}
	db.Where("user_id = ?", userID).Find(&items)
	if len(items) < 0 {
		return []*domain.UserShoppingItems{}, fmt.Errorf("not found: %w", errors.New("userShoppingItems is not found by user id"))
	}
	return items, nil
}

func (ur *UserShoppingItemRepository) FirstByID(db *gorm.DB, id int) (*domain.UserShoppingItems, error) {
	s := &domain.UserShoppingItems{}
	if err := db.First(s, id).Error; err != nil {
		return &domain.UserShoppingItems{}, fmt.Errorf("not found: %w", err)
	}
	return s, nil
}

func (ur *UserShoppingItemRepository) Create(db *gorm.DB, u *domain.UserShoppingItems) (*domain.UserShoppingItems, error) {
	if err := db.Create(u).Error; err != nil {
		return &domain.UserShoppingItems{}, fmt.Errorf("failed userShoppingItem create: %w", err)
	}
	return u, nil
}

func (sr *UserShoppingItemRepository) Save(db *gorm.DB, s *domain.UserShoppingItems) (*domain.UserShoppingItems, error) {
	if err := db.Save(s).Error; err != nil {
		return &domain.UserShoppingItems{}, fmt.Errorf("failed userShoppingItem save: %w", err)
	}
	return s, nil
}

func (ur *UserShoppingItemRepository) Delete(db *gorm.DB, id int) error {
	return db.Delete(&domain.UserShoppingItems{}, id).Error
}
