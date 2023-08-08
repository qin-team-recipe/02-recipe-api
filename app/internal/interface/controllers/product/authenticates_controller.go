package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type AuthenticatesController struct {
	Interactor product.AuthenticateInteractor
}

func NewAuthenticatesController(g gateways.Google) *AuthenticatesController {
	return &AuthenticatesController{
		Interactor: product.AuthenticateInteractor{
			Google: &gateways.GoogleGateway{Google: g},
		},
	}
}

//	@summary		GoogleアカウントログインURLの取得.
//	@description	Googleアカウントログイン認証に必要なURLの発行.
//	@tags			authenticates
//	@Success		200	{object}	controllers.H{data=product.AuthenticateResponse}
//	@Failure		400	{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/authenticates/google [get]
func (ac *AuthenticatesController) GetGoogle(ctx controllers.Context) {
	googleUrl, res := ac.Interactor.GetAuthCodeURL()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", googleUrl))
}

//	@summary		Googleアカウント情報の取得.
//	@description	Googleアカウントログイン認証に成功すればアカウント情報を取得する.
//	@tags			authenticates
//	@Param			code	query		string	true	"Googleから返却される署名（code）"
//	@Success		200		{object}	controllers.H{data=domain.SocialUserAccount}
//	@Failure		400		{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/authenticates/google/userinfo [get]
func (ac *AuthenticatesController) GetGoogleUserInfo(ctx controllers.Context) {

	code := ctx.Query("code")

	userinfo, res := ac.Interactor.GetGoogleUserInfo(code)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", userinfo))
}
