package console

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/console"
)

type RecipesController struct {
	Interactor console.RecipeInteractor
}

func NewRecipesController(db gateways.DB) *RecipesController {
	return &RecipesController{
		Interactor: console.RecipeInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Recipe: &repository.RecipeRepository{},
		},
	}
}

func (rc *RecipesController) Post(ctx controllers.Context) {
	r := &domain.Recipes{}
	if err := ctx.BindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(err.Error(), nil))
		return
	}

	recipe, res := rc.Interactor.Create(r)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipe))
}
