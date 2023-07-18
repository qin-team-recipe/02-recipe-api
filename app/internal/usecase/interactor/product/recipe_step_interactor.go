package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeStepInteractor struct {
	DB         gateway.DBRepository
	Recipe     repository.RecipeRepository
	RecipeStep repository.RecipeStepRepository
}

func (ri *RecipeStepInteractor) GetList(recipeID int) ([]*domain.RecipeStepsForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	recipeStep, err := ri.RecipeStep.FindByRecipeID(db, recipeID)
	if err != nil {
		return []*domain.RecipeStepsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipeStep, _ := ri.buildList(db, recipeStep)

	return builtRecipeStep, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeStepInteractor) Create(r *domain.RecipeSteps) (*domain.RecipeStepsForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	if _, err := ri.Recipe.FirstByID(db, r.RecipeID); err != nil {
		return &domain.RecipeStepsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	newRecipeStep, err := ri.RecipeStep.Create(db, r)
	if err != nil {
		return &domain.RecipeStepsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return newRecipeStep.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeStepInteractor) buildList(db *gorm.DB, recipeSteps []*domain.RecipeSteps) ([]*domain.RecipeStepsForGet, error) {
	builtRecipeSteps := []*domain.RecipeStepsForGet{}

	for _, recipeSteps := range recipeSteps {
		builtRecipeStep := recipeSteps.BuildForGet()

		builtRecipeSteps = append(builtRecipeSteps, builtRecipeStep)
	}

	return builtRecipeSteps, nil
}
