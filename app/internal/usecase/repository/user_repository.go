package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FirstByID(db *gorm.DB, id int) (*domain.Users, error)
	GetUser() (domain.Users, error)
	Create(db *gorm.DB, u *domain.Users) (*domain.Users, error)
	Save(db *gorm.DB, u *domain.Users) (*domain.Users, error)
	Delete(db *gorm.DB, u *domain.Users) error
}
