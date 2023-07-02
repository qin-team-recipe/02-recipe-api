package repository

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type UserOauthCertificationRepository interface {
	FirstByUserID(db *gorm.DB, userID int) (*domain.UserOauthCertifications, error)
	FirstByServiceUserID(db *gorm.DB, serviceUserID string) (*domain.UserOauthCertifications, error)
	Create(db *gorm.DB, u *domain.UserOauthCertifications) (*domain.UserOauthCertifications, error)
}
