package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type RecipeLinksController struct {
	Interactor product.RecipeLinkInteractor
}

func NewRecipeLinksController(db gateways.DB) *RecipeLinksController {
	return &RecipeLinksController{
		Interactor: product.RecipeLinkInteractor{
			DB:         &gateways.DBRepository{DB: db},
			Recipe:     &repository.RecipeRepository{},
			RecipeLink: &repository.RecipeLinkRepository{},
		},
	}
}

//	@summary		レシピURLを登録する
//	@description	レシピURLを登録し、結果を返却する
//	@tags			recipeLinks
//	@accept			application/x-json-stream
//	@param			recipe_id	formData	int		true	"Recipe ID"
//	@param			url			formData	string	true	"URL"
//	@Success		200			{object}	controllers.H{data=domain.RecipeLinksForGet}
//	@Failure		400			{object}	controllers.H{data=usecase.ResultStatus}
//	@router			/recipeLinks [post]
func (rc *RecipeLinksController) Post(ctx controllers.Context) {

	r := &domain.RecipeLinks{}

	if err := ctx.BindJSON(r); err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed bind json: %s", err.Error()), nil))
		return
	}

	recipeLink, res := rc.Interactor.Create(r)
	if res.Error != nil {
		ctx.JSON(http.StatusBadRequest, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeLink))
}
