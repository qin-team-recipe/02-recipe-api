package product

import (
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"gorm.io/gorm"
)

type ShoppingMemoInteractor struct {
	DB               gateway.DBRepository
	RecipeIngredient repository.RecipeIngredientRepository
	ShoppingMemo     repository.ShoppingMemoRepository
}

func (si *ShoppingMemoInteractor) GetList(recipeID int) ([]*domain.ShoppingMemosForGet, *usecase.ResultStatus) {

	db := si.DB.Connect()

	shoppingMemos, err := si.ShoppingMemo.FindByRecipeID(db, recipeID)
	if err != nil {
		return []*domain.ShoppingMemosForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtShoppingMemos, _ := si.buildList(db, shoppingMemos)

	return builtShoppingMemos, usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *ShoppingMemoInteractor) Create(s *domain.ShoppingMemos) (*domain.ShoppingMemosForGet, *usecase.ResultStatus) {

	db := si.DB.Connect()

	newShoppingMemo, err := si.ShoppingMemo.Create(db, s)
	if err != nil {
		return &domain.ShoppingMemosForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipeIngredient, err := si.RecipeIngredient.FirstByID(db, newShoppingMemo.RecipeIngredientID)
	if err != nil {
		return &domain.ShoppingMemosForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtShoppingMemo := newShoppingMemo.BuildForGet()
	builtShoppingMemo.RecipeIngredient = recipeIngredient.BuildForGet()

	return builtShoppingMemo, usecase.NewResultStatus(http.StatusAccepted, nil)
}

func (si *ShoppingMemoInteractor) Save(s *domain.ShoppingMemos) (*domain.ShoppingMemosForGet, *usecase.ResultStatus) {
	db := si.DB.Connect()

	foundShoppingMemo, err := si.ShoppingMemo.FirstByID(db, s.ID)

	foundShoppingMemo.IsDone = s.IsDone
	foundShoppingMemo.UpdatedAt = time.Now().Unix()

	updatedShoppingMemo, err := si.ShoppingMemo.Save(db, foundShoppingMemo)
	if err != nil {

	}
	return updatedShoppingMemo.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *ShoppingMemoInteractor) Delete(id int) *usecase.ResultStatus {

	db := si.DB.Connect()

	if _, err := si.ShoppingMemo.FirstByID(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err := si.ShoppingMemo.Delete(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *ShoppingMemoInteractor) buildList(db *gorm.DB, shoppingMemos []*domain.ShoppingMemos) ([]*domain.ShoppingMemosForGet, error) {
	builtShoppingMemos := []*domain.ShoppingMemosForGet{}

	for _, shoshoppingMemo := range shoppingMemos {
		builtShoppingMemo := shoshoppingMemo.BuildForGet()

		recipeIngredient, _ := si.RecipeIngredient.FirstByID(db, builtShoppingMemo.RecipeIngredientID)

		builtShoppingMemo.RecipeIngredient = recipeIngredient.BuildForGet()

		builtShoppingMemos = append(builtShoppingMemos, builtShoppingMemo)
	}

	return builtShoppingMemos, nil
}
