package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type RecipesController struct {
	Interactor product.RecipeInteractor
}

func NewRecipesController(db gateways.DB) *RecipesController {
	return &RecipesController{
		Interactor: product.RecipeInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Recipe: &repository.RecipeRepository{},
		},
	}
}

//	@summary		Get recipe list.
//	@description	This API return all recipe list.
//	@tags			recipes
//	@accept			application/x-json-stream
//	@Success		200	{array}		domain.RecipesForGet
//	@Failure		404	{object}	usecase.ResultStatus
//	@router			/recipes [get]
func (rc *RecipesController) GetList(ctx controllers.Context) {
	recipes, res := rc.Interactor.GetList()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipes))
}
