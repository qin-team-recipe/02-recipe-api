package product

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
)

type AuthenticateInteractor struct {
	Google gateway.GoogleGateway
}

type AuthenticateResponse struct {
	LoginURL string `json:"login_url"`
}

func (ai *AuthenticateInteractor) GetAuthCodeURL() (*AuthenticateResponse, *usecase.ResultStatus) {
	loginURL := ai.Google.AuthCodeURL("")

	if loginURL == "" {
		return &AuthenticateResponse{}, usecase.NewResultStatus(http.StatusBadGateway, errors.New("URL生成に失敗しました"))
	}

	return &AuthenticateResponse{LoginURL: loginURL}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ai *AuthenticateInteractor) GetGoogleUserInfo(code string) (*domain.Users, *usecase.ResultStatus) {
	userinfo, err := ai.Google.GetUserInfo(code)
	if err != nil {
		return &domain.Users{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	fmt.Println(userinfo)
	return &domain.Users{}, usecase.NewResultStatus(http.StatusOK, err)
}
