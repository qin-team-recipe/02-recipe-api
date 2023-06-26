package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type RecipeLinkInteractor struct {
	DB         gateway.DBRepository
	Recipe     repository.RecipeRepository
	RecipeLink repository.RecipeLinkRepository
}

func (ri *RecipeLinkInteractor) Create(r *domain.RecipeLinks) (*domain.RecipeLinksForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	if _, err := ri.Recipe.FirstByID(db, r.RecipeID); err != nil {
		return &domain.RecipeLinksForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	newRecipeLink, err := ri.RecipeLink.Create(db, r)
	if err != nil {
		return &domain.RecipeLinksForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return newRecipeLink.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}
