package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeInteractor struct {
	DB             gateway.DBRepository
	RecipeFavorite repository.RecipeFavoriteRepository
	Recipe         repository.RecipeRepository
}

// 単なるレシピのリストの取得
func (ri *RecipeInteractor) GetList(userID int, q string) ([]*domain.RecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipes, err := ri.Recipe.FindByQuery(db, userID, q)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipe, _ := ri.buildList(db, recipes)

	return builtRecipe, usecase.NewResultStatus(http.StatusOK, nil)
}

// 注目のレシピのリストを取得
func (ri *RecipeInteractor) GetRecomendList() ([]*domain.Recipes, *usecase.ResultStatus) {
	recipes := []*domain.Recipes{}

	return recipes, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeInteractor) buildList(db *gorm.DB, recipes []*domain.Recipes) ([]*domain.RecipesForGet, error) {
	builtRecipes := []*domain.RecipesForGet{}
	for _, recipe := range recipes {

		builtRecipe, err := ri.build(db, recipe)
		if err != nil {
			continue
		}

		builtRecipes = append(builtRecipes, builtRecipe)
	}

	return builtRecipes, nil
}

func (ri *RecipeInteractor) build(db *gorm.DB, recipe *domain.Recipes) (*domain.RecipesForGet, error) {
	builtRecipe := recipe.BuildForGet()

	builtRecipe.FavoritesCount = ri.RecipeFavorite.CountByRecipeID(db, builtRecipe.ID)

	return builtRecipe, nil
}
