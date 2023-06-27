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

// @Summary product users
// @Description get user info
// @Tags users,product
// @Accept application/x-json-stream
// @Success 200 {object} domain.Users
// @Router       /users/{id} [get]
func (uc *UsersController) Get(ctx controllers.Context) {
	user, res := uc.Interactor.Get()
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
	}

	ctx.JSON(res.Code, controllers.NewH("success", user))
}
