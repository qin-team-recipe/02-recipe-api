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

type RecipeStepsController struct {
	Interactor product.RecipeStepInteractor
}

func NewRecipeStepsController(db gateways.DB) *RecipeStepsController {
	return &RecipeStepsController{
		Interactor: product.RecipeStepInteractor{
			DB:         &gateways.DBRepository{DB: db},
			Recipe:     &repository.RecipeRepository{},
			RecipeStep: &repository.RecipeStepRepository{},
		},
	}
}

//	@summary		レシピの手順一覧を取得
//	@description	レシピの手順一覧を取得するエンドポイント
//	@tags			recipeSteps
//	@accept			application/x-json-stream
//	@param			recipe_id	query		int	true	"Recipe ID"
//	@Success		200			{object}	controllers.H{data=[]domain.RecipeStepsForGet}
//	@Failure		404			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeSteps [get]
func (rc *RecipeStepsController) GetList(ctx controllers.Context) {
	recipeID, _ := strconv.Atoi(ctx.Query("recipe_id"))

	Items, res := rc.Interactor.GetList(recipeID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", Items))
}


//	@summary		レシピ手順登録
//	@description	レシピ手順を登録し、結果を返却する
//	@tags			recipeSteps
//	@accept			application/x-json-stream
//	@param			recipe_id	formData	int		true	"Recipe ID"
//	@param			title		formData	string	true	"タイトル"
//	@param			description	formData	string	false	"説明"
//	@param			step_number	formData	int		false	"ステップNo"
//	@Success		200			{object}	controllers.H{data=domain.RecipeStepsForGet}
//	@Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeSteps [post]
func (rc *RecipeStepsController) Post(ctx controllers.Context) {

	r := &domain.RecipeSteps{}

	if err := ctx.BindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	recipeStep, res := rc.Interactor.Create(r)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeStep))
}
