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

type UserShoppingItemInteractor struct {
	DB               gateway.DBRepository
	UserShoppingItem repository.UserShoppingItemRepository
}

func (ui *UserShoppingItemInteractor) GetList(userID int) ([]*domain.UserShoppingItemsForGet, *usecase.ResultStatus) {
	db := ui.DB.Connect()

	userShoppingItems, err := ui.UserShoppingItem.FindByUserID(db, userID)
	if err != nil {
		return []*domain.UserShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	builtUserShoppingItems, _ := ui.buildList(db, userShoppingItems)

	return builtUserShoppingItems, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ui *UserShoppingItemInteractor) Create(u *domain.UserShoppingItems) (*domain.UserShoppingItemsForGet, *usecase.ResultStatus) {
	db := ui.DB.Connect()

	newUserShoppingItem, err := ui.UserShoppingItem.Create(db, u)
	if err != nil {
		return &domain.UserShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return newUserShoppingItem.BuildForGet(), usecase.NewResultStatus(http.StatusAccepted, nil)
}

func (si *UserShoppingItemInteractor) Save(s *domain.UserShoppingItems) (*domain.UserShoppingItemsForGet, *usecase.ResultStatus) {
	db := si.DB.Connect()

	foundUserShoppingItem, err := si.UserShoppingItem.FirstByID(db, s.ID)

	foundUserShoppingItem.Title = s.Title
	foundUserShoppingItem.Description = s.Description
	foundUserShoppingItem.IsDone = s.IsDone
	foundUserShoppingItem.UpdatedAt = time.Now().Unix()

	updatedUserShoppingItem, err := si.UserShoppingItem.Save(db, foundUserShoppingItem)
	if err != nil {
		return &domain.UserShoppingItemsForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return updatedUserShoppingItem.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (si *UserShoppingItemInteractor) Delete(id int) *usecase.ResultStatus {

	db := si.DB.Connect()

	if _, err := si.UserShoppingItem.FirstByID(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err := si.UserShoppingItem.Delete(db, id); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return usecase.NewResultStatus(http.StatusNoContent, nil)
}

func (ui *UserShoppingItemInteractor) buildList(db *gorm.DB, userShoppingItems []*domain.UserShoppingItems) ([]*domain.UserShoppingItemsForGet, error) {
	builtUserShoppingItems := []*domain.UserShoppingItemsForGet{}

	for _, userShoppingItem := range userShoppingItems {
		// builtUserShoppingItem := &domain.UserShoppingItemsForGet{}

		builtUserShoppingItems = append(builtUserShoppingItems, userShoppingItem.BuildForGet())
	}

	return builtUserShoppingItems, nil
}
