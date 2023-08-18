package product

import (
	"net/http"
	"strconv"

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
	ChefLink   repository.ChefLinkRepository
	ChefRecipe repository.ChefRecipeRepository
}

type ChefList struct {
	Lists    []*domain.ChefsForGet `json:"lists"`
	PageInfo usecase.PageInfo      `json:"page_info"`
}

func (ci *ChefInteractor) GetList(q string, cursor int) (ChefList, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chefs := []*domain.Chefs{}
	pageInfo := usecase.PageInfo{}

	// if q == "" {
	// 	foundChefs, err := ci.Chef.Find(db)
	// 	if err != nil {
	// 		return ChefList{}, usecase.NewResultStatus(http.StatusNotFound, err)
	// 	}
	// 	chefs = foundChefs
	// } else {
	// 	q = "%_" + q + "_%"
	foundChefs, err := ci.Chef.FindByQuery(db, q, cursor)
	if err != nil {
		return ChefList{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}
	chefs = foundChefs
	// }

	pageInfo.HasNextPage = 10 < len(chefs)
	pageInfo.HasPreviousPage = 0 < cursor

	if pageInfo.HasNextPage {
		pageInfo.EndCursor = strconv.Itoa(chefs[len(chefs)-1].ID)
	}

	if len(chefs) > 0 {
		pageInfo.StartCursor = strconv.Itoa(chefs[0].ID)
	}

	builtChefs, _ := ci.buildList(db, chefs)

	return ChefList{Lists: builtChefs, PageInfo: pageInfo}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ci *ChefInteractor) Get(userID int, screenName string) (*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	chef, err := ci.Chef.FirstByScreenName(db, screenName)
	if err != nil {
		return &domain.ChefsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtChef, _ := ci.build(db, chef)

	if _, err := ci.ChefFollow.FindByUserID(db, userID, 0, -1); err == nil {
		builtChef.IsFollowing = true
	}

	chefLinks, _ := ci.ChefLink.FindByChefID(db, builtChef.ID)

	builtChefLinks := []*domain.ChefLinksForGet{}
	for _, chefLink := range chefLinks {
		builtChefLinks = append(builtChefLinks, chefLink.BuildForGet())
	}

	builtChef.ChefLinks = builtChefLinks

	return builtChef, usecase.NewResultStatus(http.StatusOK, nil)
}

// おすすめのシェフリストを取得
// レコメンドの条件
func (ci *ChefInteractor) GetRecommendChefList(cursor int) ([]*domain.ChefsForGet, *usecase.ResultStatus) {

	db := ci.DB.Connect()

	// 直近三日のChefIDごとの数が多いChefFollowsを取得する
	chefFollowsCounts, err := ci.ChefFollow.FindBybyNumberOfFollowSubscriptions(db, cursor)
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
	builtChef.RecipesCount = recipeCount

	followCount := ci.ChefFollow.CountByChefID(db, chef.ID)
	builtChef.FollowsCount = followCount

	return builtChef, nil
}
