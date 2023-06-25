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

type ShoppingItemInteractor struct {
	DB               gateway.DBRepository
	RecipeIngredient repository.RecipeIngredientRepository
	ShoppingItem     repository.ShoppingItemRepository
}

func (si *ShoppingItemInteractor) GetList(recipeID int) ([]*domain.ShoppingItemsForGet, *usecase.ResultStatus) {

	db := si.DB.Connect()

	shoppingItems, err := si.ShoppingItem.FindByRecipeID(db, recipeID)
	if err != nil {
		return []*domain.ShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtShoppingItems, _ := si.buildList(db, shoppingItems)

	return builtShoppingItems, usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *ShoppingItemInteractor) Create(s *domain.ShoppingItems) (*domain.ShoppingItemsForGet, *usecase.ResultStatus) {

	db := si.DB.Connect()

	newShoppingItem, err := si.ShoppingItem.Create(db, s)
	if err != nil {
		return &domain.ShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	recipeIngredient, err := si.RecipeIngredient.FirstByID(db, newShoppingItem.RecipeIngredientID)
	if err != nil {
		return &domain.ShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	builtShoppingItem := newShoppingItem.BuildForGet()
	builtShoppingItem.RecipeIngredient = recipeIngredient.BuildForGet()

	return builtShoppingItem, usecase.NewResultStatus(http.StatusAccepted, nil)
}

func (si *ShoppingItemInteractor) Save(s *domain.ShoppingItems) (*domain.ShoppingItemsForGet, *usecase.ResultStatus) {
	db := si.DB.Connect()

	foundShoppingItem, err := si.ShoppingItem.FirstByID(db, s.ID)

	foundShoppingItem.IsDone = s.IsDone
	foundShoppingItem.UpdatedAt = time.Now().Unix()

	updatedShoppingItem, err := si.ShoppingItem.Save(db, foundShoppingItem)
	if err != nil {

	}
	return updatedShoppingItem.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *ShoppingItemInteractor) Delete(id int) *usecase.ResultStatus {

	db := si.DB.Connect()

	if _, err := si.ShoppingItem.FirstByID(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err := si.ShoppingItem.Delete(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return usecase.NewResultStatus(http.StatusNoContent, nil)
}

func (si *ShoppingItemInteractor) buildList(db *gorm.DB, shoppingItems []*domain.ShoppingItems) ([]*domain.ShoppingItemsForGet, error) {
	builtShoppingItems := []*domain.ShoppingItemsForGet{}

	for _, shoshoppingItem := range shoppingItems {
		builtShoppingItem := shoshoppingItem.BuildForGet()

		recipeIngredient, _ := si.RecipeIngredient.FirstByID(db, builtShoppingItem.RecipeIngredientID)

		builtShoppingItem.RecipeIngredient = recipeIngredient.BuildForGet()

		builtShoppingItems = append(builtShoppingItems, builtShoppingItem)
	}

	return builtShoppingItems, nil
}
