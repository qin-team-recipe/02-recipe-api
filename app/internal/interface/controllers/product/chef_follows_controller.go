package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type ChefFollowsController struct {
	Interactor product.ChefFollowInteractor
}

func NewChefFollowsController(db gateways.DB) *ChefFollowsController {
	return &ChefFollowsController{
		Interactor: product.ChefFollowInteractor{
			DB:         &gateways.DBRepository{DB: db},
			Chef:       &repository.ChefRepository{},
			ChefFollow: &repository.ChefFollowRepository{},
			User:       &repository.UserRepository{},
		},
	}
}

//	@summary		フォロー中のシェフ一覧取得
//	@description	ユーザーがフォロー中のシェフの一覧を取得する
//	@tags			chefFollows
//	@accept			application/x-json-stream
//	@param			user_id	query		int	true	"User ID"
//	@Success		200		{object}	controllers.H{data=[]domain.ChefFollowsForGet}
//	@Failure		404		{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/chefFollows [get]
func (cc *ChefFollowsController) GetList(ctx controllers.Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.Payload)

	chefFollows, res := cc.Interactor.GetList(authPayload.Audience)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chefFollows))
}

//	@summary		ユーザーがシェフをフォロー登録
//	@description	シェフをフォロー登録する際のリクエスト
//	@tags			chefFollows
//	@accept			json
//	@Param			chefFollow	body		domain.ChefFollows	true	"user_id, chef_id は必須"
//	@Success		200			{object}	controllers.H{data=domain.ChefFollowsForGet}
//	@Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/chefFollows [post]
func (cc *ChefFollowsController) Post(ctx controllers.Context) {

	f := &domain.ChefFollows{}

	if err := ctx.BindJSON(f); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	follow, res := cc.Interactor.Create(f)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", follow))
}

//	@summary		ユーザーがシェフをフォロー解除
//	@description	シェフをフォロー解除する際のリクエスト
//	@tags			chefFollows
//	@accept			json
//	@Param			chefFollow	body		domain.ChefFollows	true	"user_id, chef_id は必須"
//	@Success		200			{object}	controllers.H{data=usecase.ResultStatus}
//	@Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/chefFollows [delete]
func (cc *ChefFollowsController) Delete(ctx controllers.Context) {
	f := &domain.ChefFollows{}

	if err := ctx.BindJSON(f); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	res := cc.Interactor.Delete(f)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}

	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
