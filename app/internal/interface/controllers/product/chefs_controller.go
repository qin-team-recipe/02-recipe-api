package product

import (
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

// @summary		Get chef list.
// @description	This API return all chef list.
// @tags			chef
// @accept			application/x-json-stream
// @param			q	query		string	true	"検索ワード"
// @Success		200	{array}		domain.ChefsForGet
// @Failure		404	{object}	usecase.ResultStatus
// @router			/chefs [get]
func (cc *ChefsController) GetList(ctx controllers.Context) {
	q := ctx.Query("q")

	chefs, res := cc.Interactor.GetList(q)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chefs))
}

// @summary		Get unique chef.
// @description	This API return unique chef by screenName.
// @tags			chef
// @accept			application/x-json-stream
// @param			screenName	path		string	true	"screenName"
// @Success		200			{object}	domain.ChefsForGet
// @Failure		404			{object}	usecase.ResultStatus
// @router			/chefs/{screenName} [get]
func (cc *ChefsController) Get(ctx controllers.Context) {
	screenName := ctx.Param("screenName")

	chef, res := cc.Interactor.Get(screenName)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chef))
}
