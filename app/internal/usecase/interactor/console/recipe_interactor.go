package console

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type RecipeInteractor struct {
	DB     gateway.DBRepository
	Recipe repository.RecipeRepository
}

func (ri *RecipeInteractor) Create(recipe *domain.Recipes) (*domain.Recipes, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	newRecipe, err := ri.Recipe.Create(db, recipe)
	if err != nil {
		return &domain.Recipes{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return newRecipe, usecase.NewResultStatus(http.StatusAccepted, nil)
}
