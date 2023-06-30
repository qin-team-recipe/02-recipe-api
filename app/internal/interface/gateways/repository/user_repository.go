package repository

import (
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (ur *UserRepository) FirstByID(db *gorm.DB, id int) (*domain.Users, error) {
	user := &domain.Users{}
	if err := db.First(user, id).Error; err != nil {
		return &domain.Users{}, fmt.Errorf("user is not found: %w", err)
	}
	return user, nil
}

func (ur *UserRepository) GetUser() (domain.Users, error) {
	return domain.Users{
		ID:          100,
		ScreenName:  "FNAUHVAEM",
		DisplayName: "test taro",
	}, nil
}

func (ur *UserRepository) Create(db *gorm.DB, u *domain.Users) (*domain.Users, error) {
	if err := db.Create(u).Error; err != nil {
		return &domain.Users{}, err
	}
	return u, nil
}

func (ur *UserRepository) Save(db *gorm.DB, u *domain.Users) (*domain.Users, error) {
	if err := db.Save(u).Error; err != nil {
		return &domain.Users{}, err
	}
	return u, nil
}

func (ur *UserRepository) Delete(db *gorm.DB, u *domain.Users) error {
	if err := db.Delete(u).Error; err != nil {
		return err
	}
	return nil
}
