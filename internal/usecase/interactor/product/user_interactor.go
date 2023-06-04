package product

import (
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
)

type UserInteractor struct {
	User repository.UserRepository
}

func (ui *UserInteractor) Get() (domain.Users, *usecase.ResultStatus) {

	user, err := ui.User.GetUser()
	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(http.StatusNotFound, fmt.Errorf("test error: %w", "sample error..."))
	}
	return user, usecase.NewResultStatus(http.StatusOK, nil)
}
