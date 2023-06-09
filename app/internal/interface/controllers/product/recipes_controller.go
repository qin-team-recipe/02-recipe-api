package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qin-team-recipe/02-recipe-api/constants"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/controllers"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways"
	"github.com/qin-team-recipe/02-recipe-api/internal/interface/gateways/repository"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/interactor/product"
	"github.com/qin-team-recipe/02-recipe-api/pkg/token"
)

type RecipesController struct {
	Interactor product.RecipeInteractor
}

func NewRecipesController(db gateways.DB) *RecipesController {
	return &RecipesController{
		Interactor: product.RecipeInteractor{
			Chef:           &repository.ChefRepository{},
			ChefRecipe:     &repository.ChefRecipeRepository{},
			DB:             &gateways.DBRepository{DB: db},
			Recipe:         &repository.RecipeRepository{},
			RecipeFavorite: &repository.RecipeFavoriteRepository{},
			User:           &repository.UserRepository{},
			UserRecipe:     &repository.UserRecipeRepository{},
		},
	}
}

func (rc *RecipesController) GetList(ctx controllers.Context, jwt token.Maker) {

	authToken := ctx.GetHeader(constants.AuthorizationHeaderKey)

	userID := 0

	if authToken != "" {
		payload, err := jwt.VerifyJwtToken(authToken)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, controllers.NewH(fmt.Sprintf("failed verify jwt: %s", err.Error()), nil))
			return
		}
		userID = payload.Audience
	}

	q := ctx.Query("q")

	recipes, res := rc.Interactor.GetList(userID, q)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipes))

}

func (rc *RecipesController) Get(ctx controllers.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	recipe, res := rc.Interactor.Get(id)
	if res.Error != nil {
		ctx.JSON(res.Code, controllers.NewH(res.Error.Error(), nil))
		return
	}
	ctx.JSON(res.Code, controllers.NewH("success", recipe))
}
