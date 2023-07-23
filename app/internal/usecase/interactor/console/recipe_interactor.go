package console

import (
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"github.com/qin-team-recipe/02-recipe-api/utils"
)

type RecipeInteractor struct {
	DB         gateway.DBRepository
	ChefRecipe repository.ChefRecipeRepository
	Recipe     repository.RecipeRepository
}

func (ri *RecipeInteractor) Create(chefID int, recipe *domain.Recipes) (*domain.Recipes, *usecase.ResultStatus) {
	db := ri.DB.Begin()

	recipe.WatchID = utils.RandomString(15)
	for {
		_, err := ri.Recipe.FirstByWatchID(db, recipe.WatchID)
		if err != nil {
			break
		}
		recipe.WatchID = utils.RandomString(15)
	}

	currentTime := time.Now().Unix()
	recipe.CreatedAt = currentTime
	recipe.UpdatedAt = currentTime

	newRecipe, err := ri.Recipe.Create(db, recipe)
	if err != nil {
		db.Rollback()
		return &domain.Recipes{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	chefRecipe := &domain.ChefRecipes{
		ChefID:    chefID,
		RecipeID:  newRecipe.ID,
		CreatedAt: currentTime,
	}

	_, err = ri.ChefRecipe.Create(db, chefRecipe)
	if err != nil {
		db.Rollback()
		return &domain.Recipes{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return newRecipe, usecase.NewResultStatus(http.StatusAccepted, nil)
}
