package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type LimitedRecipeInteractor struct {
	DB     gateway.DBRepository
	Recipe repository.RecipeRepository
}

type LimitedRecipeResponse struct {
	IsLimited bool `json:"is_limited"`
}

func (li *LimitedRecipeInteractor) Save(recipeID int) (*LimitedRecipeResponse, *usecase.ResultStatus) {
	db := li.DB.Connect()

	recipe, err := li.Recipe.FirstByID(db, recipeID)
	if err != nil {
		return &LimitedRecipeResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipe.IsLimited = true

	updatedRecipe, err := li.Recipe.Save(db, recipe)
	if err != nil {
		return &LimitedRecipeResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return &LimitedRecipeResponse{
		IsLimited: updatedRecipe.IsLimited,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}
