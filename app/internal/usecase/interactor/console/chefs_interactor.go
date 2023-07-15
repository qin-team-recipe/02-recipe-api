package console

import (
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"github.com/qin-team-recipe/02-recipe-api/utils"
)

type ChefInteractor struct {
	DB   gateway.DBRepository
	Chef repository.ChefRepository
}

func (ci *ChefInteractor) Create(chef *domain.Chefs) (*domain.Chefs, *usecase.ResultStatus) {
	db := ci.DB.Connect()
	// TODO: 後で書き換える
	chef.ScreenName = utils.RandomScreenNameID(10)

	currentTime := time.Now().Unix()
	chef.CreatedAt = currentTime
	chef.UpdatedAt = currentTime

	isDuplicate, err := ci.Chef.ExistsByScreenName(db, chef.ScreenName)

	for isDuplicate {
		if isDuplicate {
			chef.ScreenName = utils.RandomScreenNameID(10)
		}
		isDuplicate, err = ci.Chef.ExistsByScreenName(db, chef.ScreenName)
		if err != nil {
			return &domain.Chefs{}, usecase.NewResultStatus(http.StatusBadRequest, err)
		}
	}

	newChef, err := ci.Chef.Create(db, chef)
	if err != nil {
		return &domain.Chefs{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newChef, usecase.NewResultStatus(http.StatusAccepted, nil)
}
