package product

import (
	"errors"
	"net/http"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/utils"
)

type AuthenticateInteractor struct {
	Google gateway.GoogleGateway
}

type AuthenticateResponse struct {
	LoginURL string `json:"login_url"`
}

func (ai *AuthenticateInteractor) GetAuthCodeURL() (*AuthenticateResponse, *usecase.ResultStatus) {
	loginURL := ai.Google.AuthCodeURL(utils.RandomToken(20))

	if loginURL == "" {
		return &AuthenticateResponse{}, usecase.NewResultStatus(http.StatusBadGateway, errors.New("URL生成に失敗しました"))
	}

	return &AuthenticateResponse{LoginURL: loginURL}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (ai *AuthenticateInteractor) GetGoogleUserInfo(code string) (*domain.SocialUserAccount, *usecase.ResultStatus) {
	googleUserID, name, email, err := ai.Google.GetUserInfo(code)
	if err != nil {
		return &domain.SocialUserAccount{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	socialType := domain.NewSocialServiceType()

	return &domain.SocialUserAccount{
		ServiceName:   socialType.Google,
		ServiceUserID: googleUserID,
		DisplayName:   name,
		Email:         email,
	}, usecase.NewResultStatus(http.StatusOK, err)
}
