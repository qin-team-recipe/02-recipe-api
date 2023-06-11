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

func (ci *ChefInteractor) GetList(q string) ([]*domain.Chefs, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefs := []*domain.Chefs{}

	if q == "" {
		foundChefs, err := ci.Chef.Find(db)
		if err != nil {
			return []*domain.Chefs{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	} else {
		q = "%" + q + "%"
		foundChefs, err := ci.Chef.FindByQuery(db, q)
		if err != nil {
			return []*domain.Chefs{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	}

	return chefs, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) Get(screenName string) (*domain.Chefs, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chef, err := ci.Chef.FirstByScreenName(db, screenName)
	if err != nil {
		return &domain.Chefs{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}
	return chef, usecase.NewResultStatus(http.StatusOK, nil)
}
