package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type LimitedRecipesController struct {
	Interactor product.LimitedRecipeInteractor
}

type LimitedRecipeRequest struct {
	RecipeID int `json:"recipe_id"`
}

func NewLimitedRecipesController(db gateways.DB) *LimitedRecipesController {
	return &LimitedRecipesController{
		Interactor: product.LimitedRecipeInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Recipe: &repository.RecipeRepository{},
		},
	}
}

func (lc *LimitedRecipesController) Patch(ctx controllers.Context) {

	l := &LimitedRecipeRequest{}

	if err := ctx.BindJSON(l); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	code, res := lc.Interactor.Save(l.RecipeID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", code))
}
