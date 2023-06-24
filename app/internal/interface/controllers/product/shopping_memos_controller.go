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

type ShoppingMemosController struct {
	Interactor product.ShoppingMemoInteractor
}

func NewShoppingMemosController(db gateways.DB) *ShoppingMemosController {
	return &ShoppingMemosController{
		Interactor: product.ShoppingMemoInteractor{
			DB:               &gateways.DBRepository{DB: db},
			RecipeIngredient: &repository.RecipeIngredientRepository{},
			ShoppingMemo:     &repository.ShoppingMemoRepository{},
		},
	}
}

func (sc *ShoppingMemosController) Post(ctx controllers.Context) {

	s := &domain.ShoppingMemos{}

	if err := ctx.BindJSON(s); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	shoppingMemo, res := sc.Interactor.Create(s)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", shoppingMemo))
}
