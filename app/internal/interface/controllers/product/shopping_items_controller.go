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

//	@summary		Get recipes shopping items.
//	@description	This API return list of recipes shopping items by Recipe ID.
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			recipe_id	query		int	true	"Recipe ID"
//	@Success		200			{array}		domain.ShoppingItemsForGet
//	@Failure		404			{object}	usecase.ResultStatus
//	@router			/shoppingItems [get]
func (sc *ShoppingItemsController) GetList(ctx controllers.Context) {
	recipeID, _ := strconv.Atoi(ctx.Query("recipe_id"))

	Items, res := sc.Interactor.GetList(recipeID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", Items))

}

//	@summary		Regist recipes shopping items.
//	@description	This API regist shopping items at recipe and return this results data.
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			user_id					formData	int		false	"User ID"
//	@param			recipe_ingredient_id	formData	int		false	"Recipe Ingredient ID"
//	@param			id_done					formData	boolean	false	"IsDone"
//	@Success		202						{object}	domain.ShoppingItemsForGet
//	@Failure		400						{object}	usecase.ResultStatus
//	@router			/shoppingItems [post]
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

//	@summary		Update state of done.
//	@description	This API update state of done at shopping items and return this results data.
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			id						path		string	true	"ID"
//	@param			user_id					formData	int		true	"User ID"
//	@param			recipe_ingredient_id	formData	int		true	"Recipe Ingredient ID"
//	@param			id_done					formData	boolean	false	"IsDone"
//	@Success		200						{object}	domain.ShoppingItemsForGet
//	@Failure		400						{object}	usecase.ResultStatus
//	@router			/shoppingItems/{id} [patch]
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

//	@summary		Delete shopping items.
//	@description	This API delete shopping items.
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			id	path		string	true	"ID"
//	@Success		200	{nil}		nil
//	@Failure		400	{object}	usecase.ResultStatus
//	@router			/shoppingItems/{id} [delete]
func (sc *ShoppingItemsController) Delete(ctx controllers.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	res := sc.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
