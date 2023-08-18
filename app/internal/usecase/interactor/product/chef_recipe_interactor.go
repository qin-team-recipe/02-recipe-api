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
	DB             gateway.DBRepository
	ChefRecipe     repository.ChefRecipeRepository
	Recipe         repository.RecipeRepository
	RecipeFavorite repository.RecipeFavoriteRepository
}

type ChefRecipeResponse struct {
	Lists    []*domain.ChefRecipesForGet `json:"lists"`
	PageInfo usecase.PageInfo            `json:"page_info"`
}

// ここはControllerで実施するべきかも
func (ri *ChefRecipeInteractor) GetList(t string, chefID, cursor int) (ChefRecipeResponse, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	res := ChefRecipeResponse{}

	result := &usecase.ResultStatus{}

	switch t {
	case "", "latest":
		// case "latest":
		return ri.getRecipesLatest(db, chefID, cursor)
	case "favorites":
		return ri.getRecipesByFavorites(db, chefID, cursor)
	}

	return res, result
}

func (ri *ChefRecipeInteractor) getRecipesLatest(db *gorm.DB, chefID, cursor int) (ChefRecipeResponse, *usecase.ResultStatus) {
	chefRecipes, err := ri.ChefRecipe.FindByChefID(db, chefID, cursor)
	if err != nil {
		return ChefRecipeResponse{
			Lists:    []*domain.ChefRecipesForGet{},
			PageInfo: usecase.PageInfo{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipeIDs := []int{}

	for _, chefRecipe := range chefRecipes {
		recipeIDs = append(recipeIDs, chefRecipe.RecipeID)
	}

	recipes, err := ri.Recipe.FindInRecipeIDs(db, recipeIDs)
	if err != nil {
		return ChefRecipeResponse{
			Lists:    []*domain.ChefRecipesForGet{},
			PageInfo: usecase.PageInfo{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipes, _ := ri.buildList(db, recipes)
	return ChefRecipeResponse{
		Lists: builtRecipes,
		PageInfo: usecase.NewPageInfo(
			10,
			len(recipes),
			cursor,
			recipes[len(recipes)-1].ID,
			recipes[0].ID,
		),
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *ChefRecipeInteractor) getRecipesByFavorites(db *gorm.DB, chefID, cursor int) (ChefRecipeResponse, *usecase.ResultStatus) {
	// ChefIDを元にお気に入り上位を取得する
	// chef_recipes, recipe_favoritesを[chef_recipe.recipe_id == recipe_favorites.recipe_id]でjoinする
	// chef_idに紐付くものをカウント計算する
	recipeFavorites, err := ri.RecipeFavorite.FindByChefRecipeIDsAndNumberOfFavoriteSubscriptions(db, chefID, cursor)
	if err != nil {
		return ChefRecipeResponse{
			Lists:    []*domain.ChefRecipesForGet{},
			PageInfo: usecase.PageInfo{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipeIDs := []int{}

	for _, recipeFavorite := range recipeFavorites {
		recipeIDs = append(recipeIDs, int(recipeFavorite))
	}

	recipes, err := ri.Recipe.FindInRecipeIDs(db, recipeIDs)
	if err != nil {
		return ChefRecipeResponse{
			Lists:    []*domain.ChefRecipesForGet{},
			PageInfo: usecase.PageInfo{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipes, _ := ri.buildList(db, recipes)

	return ChefRecipeResponse{
		Lists: builtRecipes,
		PageInfo: usecase.NewPageInfo(
			10,
			len(recipes),
			cursor,
			recipes[len(recipes)-1].ID,
			recipes[0].ID,
		),
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *ChefRecipeInteractor) buildList(db *gorm.DB, recipes []*domain.Recipes) ([]*domain.ChefRecipesForGet, error) {
	builtRecipes := []*domain.ChefRecipesForGet{}
	for _, recipe := range recipes {
		builtRecipe, err := ri.build(db, recipe)
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
