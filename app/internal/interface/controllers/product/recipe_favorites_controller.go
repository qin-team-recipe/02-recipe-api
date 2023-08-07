package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type RecipeFavoritesController struct {
	Interactor product.RecipeFavoriteInteractor
}

func NewRecipeFavoritesController(db gateways.DB) *RecipeFavoritesController {
	return &RecipeFavoritesController{
		Interactor: product.RecipeFavoriteInteractor{
			DB:             &gateways.DBRepository{DB: db},
			Chef:           &repository.ChefRepository{},
			ChefRecipe:     &repository.ChefRecipeRepository{},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
			User:           &repository.UserRepository{},
			UserRecipe:     &repository.UserRecipeRepository{},
		},
	}
}

//	@summary		ユーザーのお気に入りレシピ一覧取得
//	@description	ユーザーのお気に入りレシピ一覧を取得する際のリクエスト
//	@tags			recipeFavorites
//	@accept			application/x-json-stream
//	@param			user_id	query		int	true	"User ID"
//	@Success		200		{object}	controllers.H{data=[]domain.RecipeFavoritesForGet}
//	@Failure		404		{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeFavorites [get]
func (rc *RecipeFavoritesController) GetList(ctx controllers.Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	recipeFavorites, res := rc.Interactor.GetList(authPayload.Audience)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeFavorites))
}

//	@summary		ユーザーがレシピをお気に入り登録
//	@description	レシピをお気に入り登録する際のリクエスト
//	@tags			recipeFavorites
//	@accept			json
//	@Param			recipeFavorite	body		domain.RecipeFavorites	true	"user_id, recipe_id は必須"
//	@Success		200				{object}	controllers.H{data=domain.RecipeFavoritesForGet}
//	@Failure		400				{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeFavorites [post]
func (rc *RecipeFavoritesController) Post(ctx controllers.Context) {

	f := &domain.RecipeFavorites{}

	if err := ctx.BindJSON(f); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	favorite, res := rc.Interactor.Create(f)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", favorite))

}

//	@summary		ユーザーがレシピをお気に入り解除
//	@description	レシピをお気に入り解除する際のリクエスト
//	@tags			recipeFavorites
//	@accept			json
//	@Param			recipeFavorite	body		domain.RecipeFavorites	true	"user_id, recipe_id は必須"
//	@Success		200				{object}	controllers.H{data=usecase.ResultStatus}
//	@Failure		400				{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeFavorites [delete]
func (rc *RecipeFavoritesController) Delete(ctx controllers.Context) {
	f := &domain.RecipeFavorites{}

	if err := ctx.BindJSON(f); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	res := rc.Interactor.Delete(f)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", nil))

}
