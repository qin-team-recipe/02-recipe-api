package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type RecipeFavoritesController struct {
	Interactor product.RecipeFavoriteInteractor
}

func NewRecipeFavoritesController(db gateways.DB) *RecipeFavoritesController {
	return &RecipeFavoritesController{
		Interactor: product.RecipeFavoriteInteractor{
			DB:             &gateways.DBRepository{DB: db},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
			User:           &repository.UserRepository{},
		},
	}
}

// @summary		Get list of recipes of favorite.
// @description	This API return list of recipes of favorite.
// @tags			recipeFavorites
// @accept			application/x-json-stream
// @param			user_id	query		int	true	"User ID"
// @Success		200		{array}		domain.RecipeFavoritesForGet
// @Failure		404		{object}	usecase.ResultStatus
// @router			/recipeFavorites [get]
func (rc *RecipeFavoritesController) GetList(ctx controllers.Context) {

	userID, _ := strconv.Atoi(ctx.Query("user_id"))

	recipeFavorites, res := rc.Interactor.GetList(userID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeFavorites))
}

// @summary		ユーザーがレシピをお気に入り登録
// @description	レシピをお気に入り登録する際のリクエスト
// @tags		recipeFavorites
// @accept		json
// @Param		recipeFavorite body domain.RecipeFavorites true "user_id, recipe_id は必須"
// @Success		200			{object}	controllers.H{data=domain.RecipeFavoritesForGet}
// @Failure		400			{object}	controllers.H
// @router			/recipeFavorites [post]
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

// @summary		ユーザーがレシピをお気に入り解除
// @description	レシピをお気に入り解除する際のリクエスト
// @tags		recipeFavorites
// @accept		json
// @Param		recipeFavorite body domain.RecipeFavorites true "user_id, recipe_id は必須"
// @Success		200			{object}	controllers.H
// @Failure		400			{object}	controllers.H
// @router			/recipeFavorites [delete]
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
