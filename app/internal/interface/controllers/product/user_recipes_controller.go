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

type UserRecipesController struct {
	Interactor product.UserRecipeInteractor
}

type userRecipeRequest struct {
	Recipe            *domain.Recipes             `json:"recipe"`
	RecipeIngredients []*domain.RecipeIngredients `json:"recipe_ingredients"`
	RecipeLinks       []*domain.RecipeLinks       `json:"recipe_links"`
	RecipeSteps       []*domain.RecipeSteps       `json:"recipe_steps"`
}

func NewUserRecipesController(db gateways.DB) *UserRecipesController {
	return &UserRecipesController{
		Interactor: product.UserRecipeInteractor{
			DB:               &gateways.DBRepository{DB: db},
			Recipe:           &repository.RecipeRepository{},
			RecipeIngredient: &repository.RecipeIngredientRepository{},
			RecipeLink:       &repository.RecipeLinkRepository{},
			RecipeStep:       &repository.RecipeStepRepository{},
			UserRecipe:       &repository.UserRecipeRepository{},
		},
	}
}

func (uc *UserRecipesController) GetList(ctx controllers.Context) {
	// userRecipes, res :=
}

func (uc *UserRecipesController) Get(ctx controllers.Context) {

	// authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	// id, _ := strconv.Atoi(ctx.Param("id"))

	// userRecipe, res := uc.Interactor.Get(id)
	// if res.Error != nil {
	// 	ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
	// 	return
	// }
	// ctx.JSON(res.Code, controllers.NewH("success", userRecipe))
}

//	@summary		一般シェフレシピ登録
//	@description	一般シェフのレシピを登録し、結果を返却する
//	@tags			userRecipes
//	@accept			application/x-json-stream
//	@param			title		formData	string	true	"タイトル"
//	@param			description	formData	string	false	"説明"
//	@param			servings	formData	int		true	"対象人数"
//	@param			is_draft	formData	boolean	false	"下書きフラグ"
//	@Success		202			{object}	controllers.H{data=domain.UserRecipesForGet}
//	@Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/userRecipes [post]
func (rc *UserRecipesController) Post(ctx controllers.Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	r := &userRecipeRequest{}

	err := ctx.BindJSON(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	userRecipe, res := rc.Interactor.Create(
		authPayload.Audience,
		r.Recipe,
		r.RecipeIngredients,
		r.RecipeLinks,
		r.RecipeSteps,
	)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userRecipe))
}
