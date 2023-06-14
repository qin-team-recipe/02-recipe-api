package console

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

func (ci *ChefInteractor) Create(chef *domain.Chefs) (*domain.Chefs, *usecase.ResultStatus) {
	db := ci.DB.Connect()

	newChef, err := ci.Chef.Create(db, chef)
	if err != nil {
		return &domain.Chefs{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newChef, usecase.NewResultStatus(http.StatusAccepted, nil)
}
