package repository

import "github.com/qin-team-recipe/02-recipe-api/internal/domain"

type UserRepository interface {
	GetUser() (domain.Users, error)
}
