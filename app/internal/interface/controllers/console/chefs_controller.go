package console

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/console"
)

type ChefsController struct {
	Interactor console.ChefInteractor
}

func NewChefsController(db gateways.DB) *ChefsController {
	return &ChefsController{
		Interactor: console.ChefInteractor{
			DB:   &gateways.DBRepository{DB: db},
			Chef: &repository.ChefRepository{},
		},
	}
}

func (cc *ChefsController) Post(ctx controllers.Context) {
	c := &domain.Chefs{}
	if err := ctx.BindJSON(c); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(err.Error(), nil))
		return
	}

	chef, res := cc.Interactor.Create(c)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chef))
}
