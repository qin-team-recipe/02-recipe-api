package product

import (
	"fmt"
	"log"

	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type MeController struct {
	Jwt token.JwtMaker
}

type MeControllerProvider struct {
	DB  gateways.DB
	Jwt token.Maker
}

func NewMeController(p MeControllerProvider) *MeController {
	return &MeController{
		Jwt: token.JwtMaker{},
	}
}

func (mc *MeController) Get(ctx controllers.Context) {}

func (mc *MeController) Post(ctx controllers.Context) {
	token := mc.Jwt.CreateToken(1)

	fmt.Println(token)

	payload, err := mc.Jwt.VerifyJwtToken(token)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(payload)
}

func (mc *MeController) Patch(ctx controllers.Context) {}
