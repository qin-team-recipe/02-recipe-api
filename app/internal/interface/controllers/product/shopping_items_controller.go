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

//	@summary		買い物リスト取得
//	@description	指定されたレシピIDに紐づく買い物リストを取得する
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			recipe_id	query		int	true	"Recipe ID"
//	@Success		200			{object}	controllers.H{data=[]domain.ShoppingItemsForGet}
//	@Failure		404			{object}	controllers.H{data=usecase.ResultStatus}
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

//	@summary		買い物リストアイテム登録
//	@description	買い物リストアイテムを登録し、結果を返却する
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			user_id					formData	int		true	"User ID"
//	@param			recipe_ingredient_id	formData	int		true	"レシピ材料ID"
//	@param			id_done					formData	boolean	true	"チェック状態"
//	@Success		202						{object}	controllers.H{data=domain.ShoppingItemsForGet}
//	@Failure		400						{object}	controllers.H{data=usecase.ResultStatus}
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

//	@summary		買い物リストアイテム更新
//	@description	買い物リストアイテムのチェック状態を更新し、結果を返却する
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			id						path		string	true	"ID"
//	@param			user_id					formData	int		true	"User ID"
//	@param			recipe_ingredient_id	formData	int		true	"Recipe Ingredient ID"
//	@param			id_done					formData	boolean	false	"IsDone"
//	@Success		200						{object}	controllers.H{data=domain.ShoppingItemsForGet}
//	@Failure		400						{object}	controllers.H{data=usecase.ResultStatus}
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

//	@summary		買い物リストアイテム削除
//	@description	買い物リストアイテムを削除し、結果を返却する
//	@tags			shoppingItems
//	@accept			application/x-json-stream
//	@param			id	path		string	true	"ID"
//	@Success		200	{object}	controllers.H{data=usecase.ResultStatus}
//	@Failure		400	{object}	controllers.H{data=usecase.ResultStatus}
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
