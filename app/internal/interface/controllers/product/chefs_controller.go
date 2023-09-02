package product

import (
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type ChefsController struct {
	Jwt        gateways.Jwt
	Interactor product.ChefInteractor
}

type ChefsControllerProvider struct {
	DB  gateways.DB
	Jwt gateways.Jwt
}

func NewChefsController(p ChefsControllerProvider) *ChefsController {
	return &ChefsController{
		Jwt: &gateways.JwtGateway{Jwt: p.Jwt},
		Interactor: product.ChefInteractor{
			DB:         &gateways.DBRepository{DB: p.DB},
			Chef:       &repository.ChefRepository{},
			ChefFollow: &repository.ChefFollowRepository{},
			ChefLink:   &repository.ChefLinkRepository{},
			ChefRecipe: &repository.ChefRecipeRepository{},
		},
	}
}

// @summary		シェフ一覧取得
// @description	シェフ一覧を取得する
// @tags			chef
// @accept			application/x-json-stream
// @param			q		query		string	false	"任意のWHERE文"
// @param			cursor	query		int		true	"現在取得中の末尾カーソル"
// @Param			limit	query		string	false	"シェフの取得件数"
// @Success		200		{object}	controllers.H{data=[]product.ChefList}
// @Failure		404		{object}	controllers.H{data=usecase.ResultStatus}
// @router			/chefs [get]
func (cc *ChefsController) GetList(ctx controllers.Context) {
	q := ctx.Query("q")
	cursor, _ := strconv.Atoi(ctx.Query("cursor"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	chefs, res := cc.Interactor.GetList(q, cursor, limit)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chefs))
}

// @summary		シェフ取得
// @description	screenNameで指定されたシェフの情報を取得する
// @tags			chef
// @accept			application/x-json-stream
// @param			screenName	path		string	true	"screenName"
// @Success		200			{object}	controllers.H{data=domain.ChefsForGet}
// @Failure		404			{object}	controllers.H{data=usecase.ResultStatus}
// @router			/chefs/{screenName} [get]
func (cc *ChefsController) Get(ctx controllers.Context) {

	token := ctx.GetHeader(constants.AuthorizationHeaderKey)

	userID := 0
	if token != "" {
		authPayload, _ := cc.Jwt.VerifyJwtToken(token)
		userID = authPayload.Audience
	}

	screenName := ctx.Param("screenName")

	chef, res := cc.Interactor.Get(userID, screenName)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", chef))
}
