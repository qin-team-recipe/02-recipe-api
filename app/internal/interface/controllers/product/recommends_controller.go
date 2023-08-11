package product

import (
	"strconv"

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
			DB:         &gateways.DBRepository{DB: db},
			Chef:       &repository.ChefRepository{},
			ChefFollow: &repository.ChefFollowRepository{},
			ChefLink:   &repository.ChefLinkRepository{},
			ChefRecipe: &repository.ChefRecipeRepository{},
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

// @summary		注目のシェフ
// @description	直近3日間の獲得フォロワー数の上位10人を取得
// @tags			recommend
// @accept			application/x-json-stream
// @Success		200	{object}	controllers.H{data=[]domain.ChefsForGet}
// @Failure		404	{object}	controllers.H{data=usecase.ResultStatus}
// @router			/recommend/chefs [get]
func (rc *RecommendsController) GetRecommendChefList(ctx controllers.Context) {

	cursor, _ := strconv.Atoi(ctx.Query("cursor"))

	chefs, res := rc.ChefInteractor.GetRecommendChefList(cursor)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", chefs))
}

// @summary		話題のレシピ
// @description	過去3日間でお気に入り登録の多かったレシピを取得
// @tags			recommend
// @accept			application/x-json-stream
// @Success		200	{object}	controllers.H{data=[]domain.RecipesForGet}
// @Failure		404	{object}	controllers.H{data=usecase.ResultStatus}
// @router			/recommend/recipes [get]
func (rc *RecommendsController) GetRecommendRecipeList(ctx controllers.Context) {

	cursor, _ := strconv.Atoi(ctx.Query("cursor"))

	recipes, res := rc.RecipeInteractor.GetRecommendRecipeList(cursor)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", recipes))

}
