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

type ShoppingItemsController struct {
	Interactor product.ShoppingItemInteractor
}

func NewShoppingItemsController(db gateways.DB) *ShoppingItemsController {
	return &ShoppingItemsController{
		Interactor: product.ShoppingItemInteractor{
			DB:               &gateways.DBRepository{DB: db},
			RecipeIngredient: &repository.RecipeIngredientRepository{},
			ShoppingItem:     &repository.ShoppingItemRepository{},
		},
	}
}

func (sc *ShoppingItemsController) GetList(ctx controllers.Context) {
	recipeID, _ := strconv.Atoi(ctx.Query("recipe_id"))

	Items, res := sc.Interactor.GetList(recipeID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", Items))

}

func (sc *ShoppingItemsController) Post(ctx controllers.Context) {

	s := &domain.ShoppingItems{}

	if err := ctx.BindJSON(s); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	shoppingItem, res := sc.Interactor.Create(s)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", shoppingItem))
}

func (sc *ShoppingItemsController) Patch(ctx controllers.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	s := &domain.ShoppingItems{}

	if err := ctx.BindJSON(s); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	s.ID = id

	shoppingItem, res := sc.Interactor.Save(s)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", shoppingItem))

}

func (sc *ShoppingItemsController) Delete(ctx controllers.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	res := sc.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
