package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type ChefRecipeInteractor struct {
	DB         gateway.DBRepository
	Recipe     repository.RecipeRepository
	ChefRecipe repository.ChefRecipeRepository
}

func (ri *ChefRecipeInteractor) GetList() ([]*domain.ChefRecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipes, err := ri.Recipe.Find(db)
	if err != nil {
		return []*domain.ChefRecipesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipes, _ := ri.buildList(db, recipes)
	return builtRecipes, usecase.NewResultStatus(http.StatusOK, err)
}

func (ri *ChefRecipeInteractor) buildList(db *gorm.DB, resipes []*domain.Recipes) ([]*domain.ChefRecipesForGet, error) {
	builtRecipes := []*domain.ChefRecipesForGet{}
	for _, reresipe := range resipes {
		builtRecipe, err := ri.build(db, reresipe)
		if err != nil {
			continue
		}

		builtRecipes = append(builtRecipes, builtRecipe)
	}

	return builtRecipes, nil
}

func (ri *ChefRecipeInteractor) build(db *gorm.DB, recipe *domain.Recipes) (*domain.ChefRecipesForGet, error) {

	chefRecipe, err := ri.ChefRecipe.FirstByRecipeID(db, recipe.ID)
	if err != nil {
		return &domain.ChefRecipesForGet{}, err
	}

	builtChefRecipe := chefRecipe.BuildForGet()

	builtChefRecipe.Recipe = recipe.BuildForGet()

	return builtChefRecipe, nil
}
