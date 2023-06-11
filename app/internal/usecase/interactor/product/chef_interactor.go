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

func (ci *ChefInteractor) GetList(q string) ([]*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefs := []*domain.Chefs{}

	if q == "" {
		foundChefs, err := ci.Chef.Find(db)
		if err != nil {
			return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	} else {
		q = "%" + q + "%"
		foundChefs, err := ci.Chef.FindByQuery(db, q)
		if err != nil {
			return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	}

	builtChefs, _ := ci.buildList(chefs)

	return builtChefs, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) Get(screenName string) (*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chef, err := ci.Chef.FirstByScreenName(db, screenName)
	if err != nil {
		return &domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtChef, _ := ci.build(chef)

	return builtChef, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) buildList(chefs []*domain.Chefs) ([]*domain.ChefsForGet, error) {
	builtChefs := []*domain.ChefsForGet{}

	for _, chef := range chefs {
		builtChef, _ := ci.build(chef)

		builtChefs = append(builtChefs, builtChef)
	}

	return builtChefs, nil
}

func (ci *ChefInteractor) build(chef *domain.Chefs) (*domain.ChefsForGet, error) {
	return chef.BuildForGet(), nil
}
