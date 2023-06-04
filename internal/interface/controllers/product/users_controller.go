package product

import (
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type UsersController struct {
	Interactor product.UserInteractor
}

func NewUsersController() *UsersController {
	return &UsersController{
		Interactor: product.UserInteractor{
			User: &repository.UserRepository{},
		},
	}
}

func (uc *UsersController) Get(ctx controllers.Context) {
	user, res := uc.Interactor.Get()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
