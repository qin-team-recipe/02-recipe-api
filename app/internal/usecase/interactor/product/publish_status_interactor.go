package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain/enum"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type PublishStatusInteractor struct {
	DB     gateway.DBRepository
	Recipe repository.RecipeRepository
}

func (li *PublishStatusInteractor) Save(recipeID int, status string) *usecase.ResultStatus {
	db := li.DB.Connect()

	recipe, err := li.Recipe.FirstByID(db, recipeID)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	var s enum.Status

	switch status {
	case "public":
		s = enum.Public
	case "limited":
		s = enum.Limited
	case "private":
		s = enum.Private
	}

	recipe.PublishedStatus = s.Value()

	_, err = li.Recipe.Save(db, recipe)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return usecase.NewResultStatus(http.StatusNoContent, nil)
}
