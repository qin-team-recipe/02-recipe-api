package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type RecipeIngredientInteractor struct {
	DB               gateway.DBRepository
	RecipeIngredient repository.RecipeIngredientRepository
}

func (ri *RecipeIngredientInteractor) Create(r *domain.RecipeIngredients) (*domain.RecipeIngredientsForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	newRecipeIngredient, err := ri.RecipeIngredient.Create(db, r)
	if err != nil {
		return &domain.RecipeIngredientsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newRecipeIngredient.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}
