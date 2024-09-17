package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type UserShoppingItemsController struct {
	Interactor product.UserShoppingItemInteractor
}

func NewUserShoppingItemsController(db gateways.DB) *UserShoppingItemsController {
	return &UserShoppingItemsController{
		Interactor: product.UserShoppingItemInteractor{
			DB:               &gateways.DBRepository{DB: db},
			UserShoppingItem: &repository.UserShoppingItemRepository{},
		},
	}
}

// @summary		買い物リストアイテム一覧取得
// @description	ユーザーのIDに紐づく買い物リストを取得する
// @tags			userShoppingItems
// @accept			application/x-json-stream
// @param			user_id	query		int	true	"User ID"
// @Success		200		{object}	controllers.H{data=[]domain.UserShoppingItemsForGet}
// @Failure		404		{object}	controllers.H{data=usecase.ResultStatus}
// @router			/userShoppingItems [get]
func (uc *UserShoppingItemsController) GetList(ctx controllers.Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	userShoppingItems, res := uc.Interactor.GetList(authPayload.Audience)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItems))
}

// @summary		買い物リストアイテム登録
// @description	買い物リストにアイテムを登録し、結果を返却する
// @tags			userShoppingItems
// @accept			application/x-json-stream
// @param			user_id		formData	int		true	"User ID"
// @param			title		formData	string	true	"タイトル"
// @param			description	formData	string	false	"説明"
// @param			is_done		formData	boolean	false	"チェック状態"
// @Success		202			{object}	controllers.H{data=domain.UserShoppingItemsForGet}
// @Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
// @router			/userShoppingItems [post]
func (uc *UserShoppingItemsController) Post(ctx controllers.Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	u := &domain.UserShoppingItems{}

	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	u.UserID = authPayload.Audience

	userShoppingItem, res := uc.Interactor.Create(u)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItem))
}

// @summary		買い物リストアイテム更新
// @description	買い物リストアイテムの情報を更新し、結果を返却する
// @tags			userShoppingItems
// @accept			application/x-json-stream
// @param			id			path		string	true	"ID"
// @param			user_id		formData	int		true	"User ID"
// @param			title		formData	string	true	"タイトル"
// @param			description	formData	string	false	"説明"
// @param			is_done		formData	boolean	false	"チェック状態"
// @Success		200			{object}	controllers.H{data=domain.UserShoppingItemsForGet}
// @Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
// @router			/userShoppingItems/{id} [patch]
func (uc *UserShoppingItemsController) Patch(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	u := &domain.UserShoppingItems{}

	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	u.ID = id

	userShoppingItem, res := uc.Interactor.Save(u)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItem))
}

// @summary		買い物リストアイテム削除
// @description	買い物リストアイテムを削除する
// @tags			userShoppingItems
// @accept			application/x-json-stream
// @param			id	path		string	true	"ID"
// @Success		200	{object}	controllers.H{data=usecase.ResultStatus}
// @Failure		400	{object}	controllers.H{data=usecase.ResultStatus}
// @router			/userShoppingItems/{id} [delete]
func (uc *UserShoppingItemsController) Delete(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res := uc.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
