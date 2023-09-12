package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"github.com/qin-team-recipe/02-recipe-api/utils"
	"gorm.io/gorm"
)

type UserRecipeInteractor struct {
	DB               gateway.DBRepository
	Recipe           repository.RecipeRepository
	RecipeIngredient repository.RecipeIngredientRepository
	RecipeLink       repository.RecipeLinkRepository
	RecipeStep       repository.RecipeStepRepository
	UserRecipe       repository.UserRecipeRepository
}

// func (ri *UserRecipeInteractor) Get(id int) (*domain.UserRecipesForGet, *usecase.ResultStatus) {

// db := ri.DB.Connect()

// recipe, err := ri.Recipe.FirstByID(db, id)
// if err != nil {
// 	return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
// }

// userRecipe, err := ri.UserRecipe.FirstByRecipeID()
// }

func (ri *UserRecipeInteractor) Create(
	userID int,
	recipe *domain.Recipes,
	ingredients []*domain.RecipeIngredients,
	links []*domain.RecipeLinks,
	steps []*domain.RecipeSteps,
) (*domain.UserRecipesForGet, *usecase.ResultStatus) {

	db := ri.DB.Begin()

	currentTime := time.Now().Unix()

	// create recipe
	newRecipe, err := ri.createRecipe(db, recipe, currentTime)
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

	// create recipe ingredient
	_, _ = ri.createRecipeIngredients(db, newRecipe.ID, ingredients, currentTime)
	// if err != nil {
	// 	db.Rollback()
	// 	return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

	// create recipe link
	_, _ = ri.createRecipeLinks(db, newRecipe.ID, links, currentTime)
	// if err != nil {
	// 	db.Rollback()
	// 	return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

	// create recipe step
	_, _ = ri.createRecipeSteps(db, newRecipe.ID, steps, currentTime)
	// if err != nil {
	// 	db.Rollback()
	// 	return &domain.UserRecipesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

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

// レシピの作成
func (ri *UserRecipeInteractor) createRecipe(db *gorm.DB, recipe *domain.Recipes, currentTime int64) (*domain.Recipes, error) {
	if recipe.Title == "" {
		return &domain.Recipes{}, errors.New("タイトルが未定です")
	}

	recipe.CreatedAt = currentTime
	recipe.UpdatedAt = currentTime

	recipe.WatchID = utils.RandomString(15)
	for {
		_, err := ri.Recipe.FirstByWatchID(db, recipe.WatchID)
		if err != nil {
			break
		}
		recipe.WatchID = utils.RandomString(15)
	}

	newRecipe, err := ri.Recipe.Create(db, recipe)
	if err != nil {
		db.Rollback()
		return &domain.Recipes{}, err
	}

	return newRecipe, nil
}

func (ri *UserRecipeInteractor) createRecipeIngredients(db *gorm.DB, recipeID int, ingredients []*domain.RecipeIngredients, currentTime int64) ([]*domain.RecipeIngredients, error) {

	newRecipeIngredients := []*domain.RecipeIngredients{}
	for _, ingredient := range ingredients {
		ingredient.RecipeID = recipeID
		newRecipeIngredient, err := ri.RecipeIngredient.Create(db, ingredient)
		if err != nil {
			continue
		}

		newRecipeIngredients = append(newRecipeIngredients, newRecipeIngredient)
	}

	return newRecipeIngredients, nil
}

func (ri *UserRecipeInteractor) createRecipeLinks(db *gorm.DB, recipeID int, links []*domain.RecipeLinks, currentTime int64) ([]*domain.RecipeLinks, error) {
	newRecipeLinks := []*domain.RecipeLinks{}

	for _, link := range links {
		link.RecipeID = recipeID
		newRecipeLink, err := ri.RecipeLink.Create(db, link)
		if err != nil {
			continue
		}

		newRecipeLinks = append(newRecipeLinks, newRecipeLink)
	}
	return newRecipeLinks, nil
}

func (ri *UserRecipeInteractor) createRecipeSteps(db *gorm.DB, recipeID int, steps []*domain.RecipeSteps, currentTime int64) ([]*domain.RecipeSteps, error) {
	newRecipeSteps := []*domain.RecipeSteps{}

	for _, step := range steps {
		step.RecipeID = recipeID
		newRecipeStep, err := ri.RecipeStep.Create(db, step)
		if err != nil {
			continue
		}

		newRecipeSteps = append(newRecipeSteps, newRecipeStep)
	}
	return newRecipeSteps, nil
}
