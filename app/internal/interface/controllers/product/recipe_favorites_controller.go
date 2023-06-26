package product

import (
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
)

type RecipeFavoritesController struct {
	Interactor product.RecipeFavoriteInteractor
}

func NewRecipeFavoritesController(db gateways.DB) *RecipeFavoritesController {
	return &RecipeFavoritesController{
		Interactor: product.RecipeFavoriteInteractor{
			DB:             &gateways.DBRepository{DB: db},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
		},
	}
}

//	@summary		Get list of recipes of favorite.
//	@description	This API return list of recipes of favorite.
//	@tags			recipeFavorites
//	@accept			application/x-json-stream
//	@param			user_id	query		int	false	"User ID"
//	@Success		200		{array}		domain.RecipeFavoritesForGet
//	@Failure		404		{object}	usecase.ResultStatus
//	@router			/recipeFavorites [get]
func (rc *RecipeFavoritesController) GetList(ctx controllers.Context) {

	userID, _ := strconv.Atoi(ctx.Query("user_id"))

	recipeFavorites, res := rc.Interactor.GetList(userID)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipeFavorites))
}
