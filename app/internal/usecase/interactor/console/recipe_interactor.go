package console

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type RecipeInteractor struct {
	DB         gateway.DBRepository
	ChefRecipe repository.ChefRecipeRepository
	Recipe     repository.RecipeRepository
}

func (ri *RecipeInteractor) Create(chefID int, recipe *domain.Recipes) (*domain.Recipes, *usecase.ResultStatus) {
	db := ri.DB.Begin()

	newRecipe, err := ri.Recipe.Create(db, recipe)
	if err != nil {
		db.Rollback()
		return &domain.Recipes{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	chefRecipe := &domain.ChefRecipes{
		ChefID:   chefID,
		RecipeID: newRecipe.ID,
	}

	_, err = ri.ChefRecipe.Create(db, chefRecipe)
	if err != nil {
		db.Rollback()
		return &domain.Recipes{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return newRecipe, usecase.NewResultStatus(http.StatusAccepted, nil)
}
