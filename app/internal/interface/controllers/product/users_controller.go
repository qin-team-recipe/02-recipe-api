package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type UsersController struct {
	Interactor product.UserInteractor
}

type UsersControllerProvider struct {
	DB     gateways.DB
	Google gateways.Google
}

func NewUsersController(p *UsersControllerProvider) *UsersController {
	return &UsersController{
		Interactor: product.UserInteractor{
			Google: &gateways.GoogleGateway{Google: p.Google},
			User:   &repository.UserRepository{},
		},
	}
}

//	@summary		テストAPI
//	@description	test taro を返却
//	@tags			users
//	@accept			application/x-json-stream
//	@param			id	path		string	true	"User ID"
//	@success		200	{object}	controllers.H{data=domain.Users}
//	@failure		404	{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/users/{id} [get]
func (uc *UsersController) Get(ctx controllers.Context) {
	user, res := uc.Interactor.Get()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
