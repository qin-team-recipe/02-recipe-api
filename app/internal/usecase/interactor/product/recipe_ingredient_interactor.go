package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeIngredientInteractor struct {
	DB               gateway.DBRepository
	RecipeIngredient repository.RecipeIngredientRepository
}

func (ri *RecipeIngredientInteractor) GetList(recipeID int) ([]*domain.RecipeIngredientsForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	recipeIngredient, err := ri.RecipeIngredient.FindByRecipeID(db, recipeID)
	if err != nil {
		return []*domain.RecipeIngredientsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipeIngredient, _ := ri.buildList(db, recipeIngredient)

	return builtRecipeIngredient, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeIngredientInteractor) Create(r *domain.RecipeIngredients) (*domain.RecipeIngredientsForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	newRecipeIngredient, err := ri.RecipeIngredient.Create(db, r)
	if err != nil {
		return &domain.RecipeIngredientsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newRecipeIngredient.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeIngredientInteractor) buildList(db *gorm.DB, recipeIngredients []*domain.RecipeIngredients) ([]*domain.RecipeIngredientsForGet, error) {
	builtRecipeIngredients := []*domain.RecipeIngredientsForGet{}

	for _, recipeIngredients := range recipeIngredients {
		builtRecipeIngredient := recipeIngredients.BuildForGet()

		builtRecipeIngredients = append(builtRecipeIngredients, builtRecipeIngredient)
	}

	return builtRecipeIngredients, nil
}