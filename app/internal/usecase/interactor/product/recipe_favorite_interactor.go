package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeFavoriteInteractor struct {
	DB             gateway.DBRepository
	Recipe         repository.RecipeRepository
	RecipeFavorite repository.RecipeFavoriteRepository
}

func (ri *RecipeFavoriteInteractor) GetList(userID int) ([]*domain.RecipeFavoritesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipeFavorites, err := ri.RecipeFavorite.FindByUserID(db, userID)
	if err != nil {
		return []*domain.RecipeFavoritesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipeFavorites, _ := ri.buildList(db, recipeFavorites)

	return builtRecipeFavorites, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeFavoriteInteractor) buildList(db *gorm.DB, recipeFavorites []*domain.RecipeFavorites) ([]*domain.RecipeFavoritesForGet, error) {
	builtRecipeFavorites := []*domain.RecipeFavoritesForGet{}

	for _, recipeFavorite := range recipeFavorites {
		builtRecipeFavorite, _ := ri.build(db, recipeFavorite)

		builtRecipeFavorites = append(builtRecipeFavorites, builtRecipeFavorite)
	}

	return builtRecipeFavorites, nil
}

func (ri *RecipeFavoriteInteractor) build(db *gorm.DB, recipeFavorite *domain.RecipeFavorites) (*domain.RecipeFavoritesForGet, error) {
	builtRecipeFavorite := recipeFavorite.BuildForGet()

	recipe, err := ri.Recipe.FirstByID(db, builtRecipeFavorite.RecipeID)
	if err != nil {
		return &domain.RecipeFavoritesForGet{}, err
	}

	builtRecipeFavorite.Recipe = recipe.BuildForGet()

	return builtRecipeFavorite, nil
}
