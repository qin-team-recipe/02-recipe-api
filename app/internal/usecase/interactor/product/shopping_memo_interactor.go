package product

import (
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type ShoppingMemoInteractor struct {
	DB               gateway.DBRepository
	RecipeIngredient repository.RecipeIngredientRepository
	ShoppingMemo     repository.ShoppingMemoRepository
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

	return builtShoppingMemo, usecase.NewResultStatus(http.StatusOK, nil)
}
