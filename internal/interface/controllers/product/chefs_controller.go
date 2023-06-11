package product

import (
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type ChefsController struct {
	Interactor product.ChefInteractor
}

func NewChefsController(db gateways.DB) *ChefsController {
	return &ChefsController{
		Interactor: product.ChefInteractor{
			DB:   &gateways.DBRepository{DB: db},
			Chef: &repository.ChefRepository{},
		},
	}
}

func (cc *ChefsController) GetList(ctx controllers.Context) {
	q := ctx.Query("q")
	chefs, res := cc.Interactor.GetList(q)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chefs))
}

func (cc *ChefsController) Get(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	chef, res := cc.Interactor.Get(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chef))
}
