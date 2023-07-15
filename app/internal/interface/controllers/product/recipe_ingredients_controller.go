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

type RecipeIngredientsController struct {
	Interactor product.RecipeIngredientInteractor
}

func NewRecipeIngretientsController(db gateways.DB) *RecipeIngredientsController {
	return &RecipeIngredientsController{
		Interactor: product.RecipeIngredientInteractor{
			DB:               &gateways.DBRepository{DB: db},
			RecipeIngredient: &repository.RecipeIngredientRepository{},
		},
	}
}

//	@summary		レシピの材料一覧を取得
//	@description	レシピの材料一覧を取得するエンドポイント
//	@tags			recipeIngredients
//	@accept			application/x-json-stream
//	@param			recipe_id	query		int	true	"Recipe ID"
//	@Success		200			{array}		domain.RecipeIngredientsForGet
//	@Failure		404			{object}	usecase.ResultStatus
//	@router			/recipeIngredients [get]
func (rc *RecipeIngredientsController) GetList(ctx controllers.Context) {
	recipeID, _ := strconv.Atoi(ctx.Query("recipe_id"))

	Items, res := rc.Interactor.GetList(recipeID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", Items))
}

//	@summary		Regist recipe ingredients.
//	@description	This API regist recipe ingredients and return this results data.
//	@tags			recipeIngredients
//	@accept			application/x-json-stream
//	@param			recipe_id	formData	int		true	"RecipeID"
//	@param			name		formData	string	true	"Name"
//	@param			description	formData	string	false	"Description"
//	@Success		200			{object}	domain.RecipeIngredientsForGet
//	@Failure		400			{object}	usecase.ResultStatus
//	@router			/recipeIngredients [post]
func (rc *RecipeIngredientsController) Post(ctx controllers.Context) {
	r := &domain.RecipeIngredients{}

	if err := ctx.BindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	recipeIngredient, res := rc.Interactor.Create(r)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeIngredient))
}
