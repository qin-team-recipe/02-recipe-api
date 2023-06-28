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

//	@summary		Get users shopping items.
//	@description	This API return list of users shopping items by User ID.
//	@tags			userShoppingItems
//	@accept			application/x-json-stream
//	@param			user_id	query		int	true	"User ID"
//	@Success		200		{array}		domain.UserShoppingItemsForGet
//	@Failure		404		{object}	usecase.ResultStatus
//	@router			/userShoppingItems [get]
func (uc *UserShoppingItemsController) GetList(ctx controllers.Context) {
	userID, _ := strconv.Atoi(ctx.Query("user_id"))

	userShoppingItems, res := uc.Interactor.GetList(userID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItems))
}

//	@summary		Regist users shopping items.
//	@description	This API regist shopping items yourself and return this results data.
//	@tags			userShoppingItems
//	@accept			application/x-json-stream
//	@param			user_id		formData	int		true	"User ID"
//	@param			title		formData	string	true	"Title"
//	@param			description	formData	string	false	"Description"
//	@param			is_done		formData	boolean	false	"isDone"
//	@Success		202			{object}	domain.UserShoppingItemsForGet
//	@Failure		400			{object}	usecase.ResultStatus
//	@router			/userShoppingItems [post]
func (uc *UserShoppingItemsController) Post(ctx controllers.Context) {

	u := &domain.UserShoppingItems{}

	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	userShoppingItem, res := uc.Interactor.Create(u)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItem))
}

//	@summary		Update state of done.
//	@description	This API update state of done at user shopping items and return this results data.
//	@tags			userShoppingItems
//	@accept			application/x-json-stream
//	@param			id			path		string	true	"ID"
//	@param			user_id		formData	int		true	"User ID"
//	@param			title		formData	string	true	"Title"
//	@param			description	formData	string	false	"Description"
//	@param			is_done		formData	boolean	false	"isDone"
//	@Success		200			{object}	domain.UserShoppingItemsForGet
//	@Failure		400			{object}	usecase.ResultStatus
//	@router			/userShoppingItems/{id} [patch]
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

//	@summary		Delete shopping items.
//	@description	This API delete user shopping items.
//	@tags			userShoppingItems
//	@accept			application/x-json-stream
//	@param			id	path		string	true	"ID"
//	@Success		200	{nil}		nil
//	@Failure		400	{object}	usecase.ResultStatus
//	@router			/userShoppingItems/{id} [delete]
func (uc *UserShoppingItemsController) Delete(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res := uc.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
