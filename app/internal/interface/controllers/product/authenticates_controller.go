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

func (ac *AuthenticatesController) GetGoogle(ctx controllers.Context) {
	googleUrl, res := ac.Interactor.GetAuthCodeURL()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", googleUrl))
}
