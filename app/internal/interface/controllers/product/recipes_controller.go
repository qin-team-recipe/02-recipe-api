package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type RecipesController struct {
	Interactor product.RecipeInteractor
}

func NewRecipesController(db gateways.DB) *RecipesController {
	return &RecipesController{
		Interactor: product.RecipeInteractor{
			Chef:           &repository.ChefRepository{},
			ChefFollow:     &repository.ChefFollowRepository{},
			ChefRecipe:     &repository.ChefRecipeRepository{},
			DB:             &gateways.DBRepository{DB: db},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
			User:           &repository.UserRepository{},
			UserRecipe:     &repository.UserRecipeRepository{},
		},
	}
}

// @summary		レシピリストの取得
// @description	レシピリストを取得する
// @tags			recipes
// @Param			type	query		string	false	"type=chefFollowとすることでフォローしているシェフの情報を取得する"
// @Param			cursor	query		string	false	"取得し返している最後のレシピリストのの識別子"
// @Param			limit	query		string	false	"レシピの取得件数"
// @Success		200		{object}	controllers.H{data=product.RecipeResponse}
// @Failure		400		{object}	controllers.H{data=usecase.ResultStatus}
// @router			/recipes [get]
func (rc *RecipesController) GetList(ctx controllers.Context, jwt token.Maker) {

	ty := ctx.Query("type")
	if ty == "chefFollow" {
		rc.getLatestRecipesFromChefsFollows(ctx, jwt)
		return
	}

	authToken := ctx.GetHeader(constants.AuthorizationHeaderKey)

	userID := 0

	if authToken != "" {
		payload, err := jwt.VerifyJwtToken(authToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed verify jwt: %s", err.Error()), nil))
			return
		}
		userID = payload.Audience
	}

	q := ctx.Query("q")
	cursor, _ := strconv.Atoi(ctx.Query("cursor"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	recipes, res := rc.Interactor.GetList(userID, q, cursor, limit)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipes))
}

// // @summary		新着レシピ一覧取得
// // @description	フォロー中シェフの新着レシピリストの取得
// // @tags			recipes
// // @Param			type	query		string	true	"type=latestとすることでフォローしているシェフの新着レシピ情報を取得する"
// // @Success		200		{object}	controllers.H{data=product.RecipeResponse}
// // @Failure		400		{object}	controllers.H{data=usecase.ResultStatus}
// // @router			/recipes [get]
func (rc *RecipesController) getLatestRecipesFromChefsFollows(ctx controllers.Context, jwt token.Maker) {
	authToken := ctx.GetHeader(constants.AuthorizationHeaderKey)

	userID := 0

	if authToken != "" {
		payload, err := jwt.VerifyJwtToken(authToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed verify jwt: %s", err.Error()), nil))
			return
		}
		userID = payload.Audience
	} else {
		ctx.JSON(http.StatusForbidden, controllers.NewH("auth token not empty", nil))
		return
	}

	cursor, _ := strconv.Atoi(ctx.Query("cursor"))

	recipes, res := rc.Interactor.GetLatestRecipesFromChefsFollows(userID, cursor)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipes))
}

// @summary		レシピ情報の取得
// @description	レシピ情報を取得する
// @tags			recipes
// @Param			watch_id	path		string	true	"レシピのWatchID"
// @Success		200			{object}	controllers.H{data=domain.RecipesForGet}
// @Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
// @router			/recipes/{id} [get]
func (rc *RecipesController) Get(ctx controllers.Context, jwt token.Maker) {

	// id, _ := strconv.Atoi(ctx.Param("id"))
	authToken := ctx.GetHeader(constants.AuthorizationHeaderKey)

	userID := 0

	if authToken != "" {
		payload, _ := jwt.VerifyJwtToken(authToken)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed verify jwt: %s", err.Error()), nil))
		// 	return
		// }
		userID = payload.Audience
	}

	watchID := ctx.Param("watchID")

	recipe, res := rc.Interactor.Get(userID, watchID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipe))
}
