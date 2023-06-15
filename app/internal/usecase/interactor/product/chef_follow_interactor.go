package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type ChefFollowInteractor struct {
	DB         gateway.DBRepository
	Chef       repository.ChefRepository
	ChefFollow repository.ChefFollowRepository
}

func (ci *ChefFollowInteractor) GetList(userID int) ([]*domain.ChefFollowsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefFollows, err := ci.ChefFollow.FindByUserID(db, userID)
	if err != nil {
		return []*domain.ChefFollowsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtChefFollows, _ := ci.buildList(db, chefFollows)

	return builtChefFollows, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefFollowInteractor) buildList(db *gorm.DB, chefFollows []*domain.ChefFollows) ([]*domain.ChefFollowsForGet, error) {

	builtChefFollows := []*domain.ChefFollowsForGet{}

	for _, chefFollow := range chefFollows {
		builtChefFollow, _ := ci.build(db, chefFollow)

		builtChefFollows = append(builtChefFollows, builtChefFollow)
	}

	return builtChefFollows, nil
}

func (ci *ChefFollowInteractor) build(db *gorm.DB, chefFollow *domain.ChefFollows) (*domain.ChefFollowsForGet, error) {
	chef, err := ci.Chef.FirstByID(db, chefFollow.ChefID)
	if err != nil {
		return &domain.ChefFollowsForGet{}, err
	}

	builtChefFollow := chefFollow.BuildForGet()

	builtChefFollow.Chef = chef.BuildForGet()

	return builtChefFollow, nil
}
