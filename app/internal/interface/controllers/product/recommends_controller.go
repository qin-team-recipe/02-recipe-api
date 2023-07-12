package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type RecommendsController struct {
	ChefInteractor   product.ChefInteractor
	RecipeInteractor product.RecipeInteractor
}

func NewRecommendsController(db gateways.DB) *RecommendsController {
	return &RecommendsController{
		ChefInteractor: product.ChefInteractor{
			DB:   &gateways.DBRepository{DB: db},
			Chef: &repository.ChefRepository{},
		},
		RecipeInteractor: product.RecipeInteractor{
			Chef:           &repository.ChefRepository{},
			ChefRecipe:     &repository.ChefRecipeRepository{},
			DB:             &gateways.DBRepository{DB: db},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
			User:           &repository.UserRepository{},
			UserRecipe:     &repository.UserRecipeRepository{},
		},
	}
}

func (rc *RecommendsController) GetRecommendChefList(ctx controllers.Context) {

	chefs, res := rc.ChefInteractor.GetRecommendChefList()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", chefs))
}

func (rc *RecommendsController) GetRecommendRecipeList(ctx controllers.Context) {

	recipes, res := rc.RecipeInteractor.GetRecommendRecipeList()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", recipes))

}
