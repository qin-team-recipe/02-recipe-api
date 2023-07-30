package product

import (
	"errors"
	"net/http"
	"sort"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeInteractor struct {
	Chef           repository.ChefRepository
	ChefFollow     repository.ChefFollowRepository
	ChefRecipe     repository.ChefRecipeRepository
	DB             gateway.DBRepository
	RecipeFavorite repository.RecipeFavoriteRepository
	Recipe         repository.RecipeRepository
	User           repository.UserRepository
	UserRecipe     repository.UserRecipeRepository
}

// 単なるレシピのリストの取得
func (ri *RecipeInteractor) GetList(userID int, q string) ([]*domain.RecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipes, err := ri.Recipe.FindByQuery(db, userID, q)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipes, _ := ri.buildList(db, recipes)

	return builtRecipes, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeInteractor) GetLatestRecipesFromChefsFollows(userID int) ([]*domain.RecipesForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	chefFollows, err := ri.ChefFollow.FindByUserID(db, userID)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	followIDs := []int{}
	for _, chefFollow := range chefFollows {
		followIDs = append(followIDs, chefFollow.ChefID)
	}

	chefRecipes, err := ri.ChefRecipe.FindInByChefIDs(db, followIDs)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipeIDs := []int{}
	for _, chefRecipe := range chefRecipes {
		recipeIDs = append(followIDs, chefRecipe.RecipeID)
	}

	recipes, err := ri.Recipe.FindInRecipeIDs(db, recipeIDs)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	sort.Slice(recipes, func(i, j int) bool { return recipes[i].CreatedAt > recipes[j].CreatedAt })

	builtRecipes, _ := ri.buildList(db, recipes)
	return builtRecipes, usecase.NewResultStatus(http.StatusOK, nil)
}

// 注目のレシピのリストを取得
// レコメンドの条件
func (ri *RecipeInteractor) GetRecommendRecipeList() ([]*domain.RecipesForGet, *usecase.ResultStatus) {

	db := ri.DB.Connect()

	// 直近三日のRecipeIDごとの数が多いRecipeIDとCountを取得する
	recipeFavoritesCounts, err := ri.RecipeFavorite.FindByNumberOfFavoriteSubscriptions(db)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("注目されているレシピはありません"))
	}

	recipeIDs := []int{}
	for key := range recipeFavoritesCounts {
		recipeIDs = append(recipeIDs, key)
	}

	recipes, err := ri.Recipe.FindInRecipeIDs(db, recipeIDs)
	if err != nil {
		return []*domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipes, _ := ri.buildList(db, recipes)

	return builtRecipes, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeInteractor) Get(watchID string) (*domain.RecipesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	recipe, err := ri.Recipe.FirstByWatchID(db, watchID)
	if err != nil {
		return &domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtRecipe, err := ri.build(db, recipe)
	if err != nil {
		return &domain.RecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return builtRecipe, usecase.NewResultStatus(http.StatusOK, nil)
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

	if builtRecipe.IsDraft {
		return &domain.RecipesForGet{}, errors.New("下書き状態のレシピです")
	}

	builtRecipe.FavoritesCount = ri.RecipeFavorite.CountByRecipeID(db, builtRecipe.ID)

	chefRecipe, err := ri.ChefRecipe.FirstByRecipeID(db, builtRecipe.ID)
	if err == nil {
		chef, _ := ri.Chef.FirstByID(db, chefRecipe.ChefID)

		builtRecipe.Chef = chef.BuildForGet()
	} else {
		userRecipe, err := ri.UserRecipe.FirstByRecipeID(db, builtRecipe.ID)
		if err == nil {
			user, err := ri.User.FirstByID(db, userRecipe.UserID)
			if err != nil {
				return &domain.RecipesForGet{}, errors.New("レシピを作成したシェフ、またはユーザーが見つかりません")
			}

			builtRecipe.User = user.BuildForGet()
		}
	}

	return builtRecipe, nil
}
