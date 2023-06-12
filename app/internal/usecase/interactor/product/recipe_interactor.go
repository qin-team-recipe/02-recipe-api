package product

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

func (ri *RecipeInteractor) GetList() ([]*domain.RecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipes, err := ri.Recipe.Find(db)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipes, _ := ri.buildList(recipes)
	return builtRecipes, usecase.NewResultStatus(http.StatusOK, err)
}

func (ri *RecipeInteractor) buildList(resipes []*domain.Recipes) ([]*domain.RecipesForGet, error) {
	builtRecipes := []*domain.RecipesForGet{}
	for _, reresipe := range resipes {
		builtRecipe, _ := ri.build(reresipe)

		builtRecipes = append(builtRecipes, builtRecipe)
	}

	return builtRecipes, nil
}

func (ri *RecipeInteractor) build(recipe *domain.Recipes) (*domain.RecipesForGet, error) {
	return recipe.BuildForGet(), nil
}
