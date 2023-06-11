package repository

import "github.com/qin-team-recipe/02-recipe-api/internal/domain"

type UserRepository struct{}

func (ur *UserRepository) GetUser() (domain.Users, error) {
	return domain.Users{
		ID:          100,
		ScreenName:  "FNAUHVAEM",
		DisplayName: "test taro",
	}, nil
}
