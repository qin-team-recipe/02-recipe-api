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

func (uc *UserShoppingItemsController) GetList(ctx controllers.Context) {
	userID, _ := strconv.Atoi(ctx.Query("user_id"))

	userShoppingItems, res := uc.Interactor.GetList(userID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userShoppingItems))
}

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

func (uc *UserShoppingItemsController) Delete(ctx controllers.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res := uc.Interactor.Delete(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
