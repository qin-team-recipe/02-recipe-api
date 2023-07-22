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

func (li *LimitedRecipeInteractor) Save(recipeID int) *usecase.ResultStatus {
	db := li.DB.Connect()

	recipe, err := li.Recipe.FirstByID(db, recipeID)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	switch recipe.IsLimited {
	case true:
		recipe.IsLimited = false
	case false:
		recipe.IsLimited = true
	}

	_, err = li.Recipe.Save(db, recipe)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return usecase.NewResultStatus(http.StatusNoContent, nil)
}
