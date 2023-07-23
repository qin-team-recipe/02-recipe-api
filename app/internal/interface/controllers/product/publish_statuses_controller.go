package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type PublichStatusesController struct {
	Interactor product.PublishStatusInteractor
}

type PublishStatusRequest struct {
	RecipeID int    `json:"recipe_id"`
	Status   string `json:"status"`
}

func NewPublishStatusesController(db gateways.DB) *PublichStatusesController {
	return &PublichStatusesController{
		Interactor: product.PublishStatusInteractor{
			DB:     &gateways.DBRepository{DB: db},
			Recipe: &repository.RecipeRepository{},
		},
	}
}

// @summary		レシピの非公開状態にする
// @description	レシピを非公開状態にする
// @tags			recipes
// @Param		publish_status_reqest	body	product.PublishStatusRequest		true	"レシピのIDとステータスを含む"
// @Success		200		{object}	controllers.H
// @Failure		400		{object}	controllers.H
// @router			/publishStatuses [patch]
func (lc *PublichStatusesController) Patch(ctx controllers.Context) {

	l := &PublishStatusRequest{}

	if err := ctx.BindJSON(l); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	res := lc.Interactor.Save(l.RecipeID, l.Status)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
