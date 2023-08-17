package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type MeController struct {
	Auth       product.AuthenticateInteractor
	Interactor product.MeInteractor
}

type MeControllerProvider struct {
	DB     gateways.DB
	Google gateways.Google
	Jwt    gateways.Jwt
}

func NewMeController(p MeControllerProvider) *MeController {
	return &MeController{
		Auth: product.AuthenticateInteractor{
			Google: &gateways.GoogleGateway{Google: p.Google},
		},
		Interactor: product.MeInteractor{
			DB:                     &gateways.DBRepository{DB: p.DB},
			Jwt:                    &gateways.JwtGateway{Jwt: p.Jwt},
			User:                   &repository.UserRepository{},
			UserOauthCertification: &repository.UserOauthCertificationRepository{},
		},
	}
}

// @summary		ユーザーログイン
// @description	ユーザーログインのエンドポイント
// @tags			me
// @Param			service_user_id	query		string	true	"GoogleアカウントのユーザーID"
// @Success		200				{object}	controllers.H{data=product.UserResponse}
// @Failure		400				{object}	controllers.H{data=usecase.ResultStatus}
// @router			/login [get]
func (mc *MeController) LoginUser(ctx controllers.Context) {

	serviceUserID := ctx.Query("service_user_id")

	user, res := mc.Interactor.LoginUser(serviceUserID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", user))
}

// @summary		ユーザー情報の取得
// @description	ユーザー情報の取得のエンドポイント
// @tags			me
// @Success		200	{object}	controllers.H{data=domain.Users}
// @Failure		400	{object}	controllers.H{data=usecase.ResultStatus}
// @router			/me [get]
func (mc *MeController) Get(ctx controllers.Context) {
	authPayload := ctx.MustGet("authorization_payload").(*token.Payload)
	// authToken := ctx.GetHeader("authorization")

	// userID := authPayload.Audience

	me, res := mc.Interactor.Get(authPayload.Audience)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", me.BuildForGet()))
}

// @summary		ユーザー新規登録
// @description	ユーザー新規登録のエンドポイント
// @tags			me
// @Param			social_user_account	body		domain.SocialUserAccount	true	"登録するGoogleアカウント"
// @Success		200					{object}	controllers.H{data=product.UserResponse}
// @Failure		400					{object}	controllers.H{data=usecase.ResultStatus}
// @router			/register [post]
func (mc *MeController) Post(ctx controllers.Context) {

	u := &domain.SocialUserAccount{}
	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	me, res := mc.Interactor.Create(u)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", me))
}

// @summary		ユーザー情報の変更
// @description	ユーザー情報の変更のエンドポイント
// @tags			me
// @Param			user	body		domain.Users	true	"変更したユーザー情報"
// @Success		200		{object}	controllers.H{data=domain.UsersForGet}
// @Failure		400		{object}	controllers.H{data=usecase.ResultStatus}
// @router			/me [patch]
func (mc *MeController) Patch(ctx controllers.Context) {

	u := &domain.Users{}

	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	me, res := mc.Interactor.Save(u)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", me))
}

// @summary		ユーザーアカウントの削除
// @description	ユーザーアカウントの削除のエンドポイント
// @tags			me
// @Success		200	{object}	controllers.H{data=usecase.ResultStatus}
// @Failure		400	{object}	controllers.H{data=usecase.ResultStatus}
// @router			/me [delete]
func (mc *MeController) Delete(ctx controllers.Context) {

	authToken := ctx.GetHeader("authorization")

	res := mc.Interactor.Delete(authToken)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
