package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserShoppingItemRepository interface {
	FindByUserID(db *gorm.DB, userID int) ([]*domain.UserShoppingItems, error)
	FirstByID(db *gorm.DB, id int) (*domain.UserShoppingItems, error)
	Create(db *gorm.DB, u *domain.UserShoppingItems) (*domain.UserShoppingItems, error)
	Save(db *gorm.DB, u *domain.UserShoppingItems) (*domain.UserShoppingItems, error)
	Delete(db *gorm.DB, id int) error
}
