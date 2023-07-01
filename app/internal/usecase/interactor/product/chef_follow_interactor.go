package product

import (
	"errors"
	"net/http"
	"time"

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
	User       repository.UserRepository
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

func (ci *ChefFollowInteractor) Create(chefFollow *domain.ChefFollows) (*domain.ChefFollowsForGet, *usecase.ResultStatus) {
	db := ci.DB.Connect()

	if _, err := ci.User.FirstByID(db, chefFollow.UserID); err != nil {
		return &domain.ChefFollowsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if _, err := ci.Chef.FirstByID(db, chefFollow.ChefID); err != nil {
		return &domain.ChefFollowsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if _, err := ci.ChefFollow.FirstByUserIDAndChefID(db, chefFollow.UserID, chefFollow.ChefID); err == nil {
		return &domain.ChefFollowsForGet{}, usecase.NewResultStatus(http.StatusConflict, errors.New("既にフォローしています"))
	}

	chefFollow.CreatedAt = time.Now().Unix()

	newChefFollow, err := ci.ChefFollow.Create(db, chefFollow)
	if err != nil {
		return &domain.ChefFollowsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newChefFollow.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefFollowInteractor) Delete(f *domain.ChefFollows) *usecase.ResultStatus {

	db := ci.DB.Connect()

	follow, err := ci.ChefFollow.FirstByUserIDAndChefID(db, f.UserID, f.ChefID)
	if err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err = ci.ChefFollow.Delete(db, follow); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return usecase.NewResultStatus(http.StatusOK, nil)
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
