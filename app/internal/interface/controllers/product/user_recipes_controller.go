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
			DB:         &gateways.DBRepository{DB: db},
			Recipe:     &repository.RecipeRepository{},
			UserRecipe: &repository.UserRecipeRepository{},
		},
	}
}

func (uc *UserRecipesController) GetList(ctx controllers.Context) {
	// userRecipes, res :=
}

// @summary		Regist user recipes.
// @description	This API regist user recipes and return this results data.
// @tags			userRecipes
// @accept			application/x-json-stream
// @param			title		formData	string	true	"Title"
// @param			description	formData	string	false	"Description"
// @param			servings	formData	int		true	"Servings"
// @param			is_draft	formData	boolean	false	"isDraft"
// @Success		202			{object}	domain.UserRecipesForGet
// @Failure		400			{object}	usecase.ResultStatus
// @router			/userRecipes [post]
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
