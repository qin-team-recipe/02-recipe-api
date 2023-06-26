package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type UserRecipesController struct {
	Interactor product.UserRecipeInteractor
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

func (rc *UserRecipesController) Post(ctx controllers.Context) {
	// userID(仮)
	userID := 1

	r := &domain.Recipes{}

	err := ctx.BindJSON(r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	userRecipe, res := rc.Interactor.Create(userID, r)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userRecipe))
}
