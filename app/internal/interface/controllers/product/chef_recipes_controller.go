package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type ChefRecipesController struct {
	Interactor product.ChefRecipeInteractor
}

func NewChefRecipesController(db gateways.DB) *ChefRecipesController {
	return &ChefRecipesController{
		Interactor: product.ChefRecipeInteractor{
			DB:         &gateways.DBRepository{DB: db},
			Recipe:     &repository.RecipeRepository{},
			ChefRecipe: &repository.ChefRecipeRepository{},
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
func (rc *ChefRecipesController) GetList(ctx controllers.Context) {
	recipes, res := rc.Interactor.GetList()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipes))
}
