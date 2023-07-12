package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type ChefInteractor struct {
	DB         gateway.DBRepository
	Chef       repository.ChefRepository
	ChefFollow repository.ChefFollowRepository
	ChefRecipe repository.ChefRecipeRepository
}

func (ci *ChefInteractor) GetList(q string) ([]*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefs := []*domain.Chefs{}

	if q == "" {
		foundChefs, err := ci.Chef.Find(db)
		if err != nil {
			return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	} else {
		q = "%" + q + "%"
		foundChefs, err := ci.Chef.FindByQuery(db, q)
		if err != nil {
			return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
		}
		chefs = foundChefs
	}

	builtChefs, _ := ci.buildList(db, chefs)

	return builtChefs, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) Get(userID int, screenName string) (*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chef, err := ci.Chef.FirstByScreenName(db, screenName)
	if err != nil {
		return &domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtChef, _ := ci.build(db, chef)

	if _, err := ci.ChefFollow.FindByUserID(db, userID); err == nil {
		builtChef.IsFollowing = true
	}

	return builtChef, usecase.NewResultStatus(http.StatusOK, nil)
}

// おすすめのシェフリストを取得
// レコメンドの条件
func (ci *ChefInteractor) GetRecommendChefList() ([]*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	// 直近三日のChefIDごとの数が多いChefFollowsを取得する
	chefFollowsCounts, err := ci.ChefFollow.FindBybyNumberOfFollowSubscriptions(db)
	if err != nil {
		return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	chefIDs := []int{}
	for key := range chefFollowsCounts {
		chefIDs = append(chefIDs, key)
	}

	chefs, err := ci.Chef.FindInChefIDs(db, chefIDs)
	if err != nil {
		return []*domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtChefs, _ := ci.buildList(db, chefs)

	return builtChefs, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) buildList(db *gorm.DB, chefs []*domain.Chefs) ([]*domain.ChefsForGet, error) {
	builtChefs := []*domain.ChefsForGet{}

	for _, chef := range chefs {
		builtChef, _ := ci.build(db, chef)

		builtChefs = append(builtChefs, builtChef)
	}

	return builtChefs, nil
}

func (ci *ChefInteractor) build(db *gorm.DB, chef *domain.Chefs) (*domain.ChefsForGet, error) {

	builtChef := &domain.ChefsForGet{}

	builtChef = chef.BuildForGet()

	recipeCount := ci.ChefRecipe.CountByChefID(db, chef.ID)
	builtChef.RecipeCount = recipeCount

	followCount := ci.ChefFollow.CountByChefID(db, chef.ID)
	builtChef.FollowsCount = followCount

	return builtChef, nil
}
