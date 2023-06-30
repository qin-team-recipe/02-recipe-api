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

func (mc *MeController) LoginUser(ctx controllers.Context) {

	serviceUserID := ctx.Query("service_user_id")

	user, res := mc.Interactor.LoginUser(serviceUserID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", user))
}

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

func (mc *MeController) Delete(ctx controllers.Context) {

	authToken := ctx.GetHeader("authorization")

	res := mc.Interactor.Delete(authToken)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", nil))
}
