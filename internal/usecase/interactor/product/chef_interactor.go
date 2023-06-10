package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type ChefInteractor struct {
	DB   gateway.DBRepository
	Chef repository.ChefRepository
}

func (ci *ChefInteractor) GetList() ([]*domain.Chefs, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefs, err := ci.Chef.Find(db)
	if err != nil {
		return []*domain.Chefs{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	return chefs, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) Get(id int) (*domain.Chefs, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chef, err := ci.Chef.FirstByID(db, id)
	if err != nil {
		return &domain.Chefs{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}
	return chef, usecase.NewResultStatus(http.StatusOK, nil)
}
