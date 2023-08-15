package product

import (
	"errors"
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type RecipeFavoriteInteractor struct {
	DB             gateway.DBRepository
	Chef           repository.ChefRepository
	ChefRecipe     repository.ChefRecipeRepository
	Recipe         repository.RecipeRepository
	RecipeFavorite repository.RecipeFavoriteRepository
	User           repository.UserRepository
	UserRecipe     repository.UserRecipeRepository
}

type RecipeFavoriteResponse struct {
	Lists    []*domain.RecipeFavoritesForGet `json:"lists"`
	PageInfo usecase.PageInfo                `json:"page_info"`
}

func (ri *RecipeFavoriteInteractor) GetList(userID, cursor, limit int) (RecipeFavoriteResponse, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	if limit <= 0 {
		limit = 10
	}

	recipeFavorites, err := ri.RecipeFavorite.FindByUserID(db, userID, cursor, limit+1)
	if err != nil {
		return RecipeFavoriteResponse{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtRecipeFavorites, _ := ri.buildList(db, recipeFavorites)

	return RecipeFavoriteResponse{
		Lists:    builtRecipeFavorites,
		PageInfo: usecase.NewPageInfo(10, len(builtRecipeFavorites), cursor, builtRecipeFavorites[len(builtRecipeFavorites)-1].ID, builtRecipeFavorites[0].ID),
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeFavoriteInteractor) Create(f *domain.RecipeFavorites) (*domain.RecipeFavoritesForGet, *usecase.ResultStatus) {
	db := ri.DB.Connect()

	if _, err := ri.User.FirstByID(db, f.UserID); err != nil {
		return &domain.RecipeFavoritesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if _, err := ri.Recipe.FirstByID(db, f.RecipeID); err != nil {
		return &domain.RecipeFavoritesForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if _, err := ri.RecipeFavorite.FirstByUserIDAndRecipeID(db, f.UserID, f.RecipeID); err == nil {
		return &domain.RecipeFavoritesForGet{}, usecase.NewResultStatus(http.StatusConflict, errors.New("既にお気に入りに登録されています。"))
	}

	f.CreatedAt = time.Now().Unix()

	newRecipeFavorite, err := ri.RecipeFavorite.Create(db, f)
	if err != nil {
		return &domain.RecipeFavoritesForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return newRecipeFavorite.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeFavoriteInteractor) Delete(f *domain.RecipeFavorites) *usecase.ResultStatus {

	db := ri.DB.Connect()

	favorite, err := ri.RecipeFavorite.FirstByUserIDAndRecipeID(db, f.UserID, f.RecipeID)
	if err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err = ri.RecipeFavorite.Delete(db, favorite); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return usecase.NewResultStatus(http.StatusOK, nil)
}

func (ri *RecipeFavoriteInteractor) buildList(db *gorm.DB, recipeFavorites []*domain.RecipeFavorites) ([]*domain.RecipeFavoritesForGet, error) {
	builtRecipeFavorites := []*domain.RecipeFavoritesForGet{}

	for _, recipeFavorite := range recipeFavorites {
		builtRecipeFavorite, err := ri.build(db, recipeFavorite)
		if err != nil {
			continue
		}

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

	chefRecipe, err := ri.ChefRecipe.FirstByRecipeID(db, recipe.ID)
	if err == nil {
		chef, _ := ri.Chef.FirstByID(db, chefRecipe.ChefID)
		builtRecipeFavorite.Recipe.Chef = chef.BuildForGet()
	} else {
		userRecipe, err := ri.UserRecipe.FirstByRecipeID(db, recipe.ID)
		if err != nil {
			return &domain.RecipeFavoritesForGet{}, err
		}

		user, err := ri.User.FirstByID(db, userRecipe.UserID)
		if err != nil {
			return &domain.RecipeFavoritesForGet{}, err
		}
		builtRecipeFavorite.Recipe.User = user.BuildForGet()
	}

	builtRecipeFavorite.Recipe.FavoritesCount = ri.RecipeFavorite.CountByRecipeID(db, recipe.ID)

	return builtRecipeFavorite, nil
}
