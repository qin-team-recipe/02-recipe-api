package product

import (
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type UserRecipeInteractor struct {
	DB         gateway.DBRepository
	Recipe     repository.RecipeRepository
	UserRecipe repository.UserRecipeRepository
}

func (ri *UserRecipeInteractor) Create(userID int, r *domain.Recipes) (*domain.UserRecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Begin()

	currentTime := time.Now().Unix()

	r.CreatedAt = currentTime
	r.UpdatedAt = currentTime

	newRecipe, err := ri.Recipe.Create(db, r)
	if err != nil {
		db.Rollback()
		return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	userRecipe, err := ri.UserRecipe.Create(db, &domain.UserRecipes{
		UserID:    userID,
		RecipeID:  newRecipe.ID,
		CreatedAt: currentTime,
	})
	if err != nil {
		db.Rollback()
		return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtUserRecipe, err := ri.build(userRecipe, newRecipe)
	if err != nil {
		db.Rollback()
		return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return builtUserRecipe, usecase.NewResultStatus(http.StatusAccepted, nil)
}

func (ri *UserRecipeInteractor) buildList(db *gorm.DB, userResipes []*domain.UserRecipes) ([]*domain.UserRecipesForGet, error) {
	builtUserRecipes := []*domain.UserRecipesForGet{}
	for _, userRecipe := range userResipes {
		recipe, _ := ri.Recipe.FirstByID(db, userRecipe.RecipeID)
		builtUserRecipe, _ := ri.build(userRecipe, recipe)

		builtUserRecipes = append(builtUserRecipes, builtUserRecipe)
	}

	return builtUserRecipes, nil
}

func (ri *UserRecipeInteractor) build(userRecipe *domain.UserRecipes, recipe *domain.Recipes) (*domain.UserRecipesForGet, error) {
	builtUserRecipe := userRecipe.BuildForGet()

	builtUserRecipe.Recipe = recipe.BuildForGet()
	return builtUserRecipe, nil
}
