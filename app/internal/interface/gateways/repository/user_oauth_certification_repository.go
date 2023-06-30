package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserOauthCertificationRepository struct{}

func (ur *UserOauthCertificationRepository) FirstByUserID(db *gorm.DB, userID int) (*domain.UserOauthCertifications, error) {
	u := &domain.UserOauthCertifications{}
	if err := db.Where("user_id = ?", userID).First(u).Error; err != nil {
		return &domain.UserOauthCertifications{}, err
	}
	return u, nil
}

func (ur *UserOauthCertificationRepository) FirstByServiceUserID(db *gorm.DB, serviceUserID string) (*domain.UserOauthCertifications, error) {
	u := &domain.UserOauthCertifications{}
	if err := db.Where("service_user_id = ?", serviceUserID).First(u).Error; err != nil {
		return &domain.UserOauthCertifications{}, err
	}
	return u, nil
}

func (ur *UserOauthCertificationRepository) Create(db *gorm.DB, u *domain.UserOauthCertifications) (*domain.UserOauthCertifications, error) {
	if err := db.Create(u).Error; err != nil {
		return &domain.UserOauthCertifications{}, err
	}
	return u, nil
}
