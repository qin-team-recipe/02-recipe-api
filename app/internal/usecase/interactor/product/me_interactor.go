package product

import (
	"fmt"
	"net/http"
	"time"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/gateway"
	"github.com/qin-team-recipe/02-recipe-api/internal/usecase/repository"
	"github.com/qin-team-recipe/02-recipe-api/pkg/random"
)

type MeInteractor struct {
	DB                     gateway.DBRepository
	Jwt                    gateway.JwtGateway
	Redis                  gateway.RedisGateway
	User                   repository.UserRepository
	UserOauthCertification repository.UserOauthCertificationRepository
}

type UserResponse struct {
	User  *domain.UsersForGet `json:"user"`
	Token string              `json:"token"`
}

func (mi *MeInteractor) LoginUser(serviceUserID int) (UserResponse, *usecase.ResultStatus) {
	db := mi.DB.Connect()

	userOauthCretificate, err := mi.UserOauthCertification.FirstByUserID(db, serviceUserID)
	if err != nil {
		return UserResponse{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	user, err := mi.User.FirstByID(db, userOauthCretificate.UserID)
	if err != nil {
		return UserResponse{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	jwtToken, payload, err := mi.Jwt.CreateToken(user.ID)
	if err != nil {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	fmt.Println(payload)
	// err = mi.Redis.Set(payload.ID, payload)
	// if err != nil {
	// 	return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

	return UserResponse{
		User:  user.BuildForGet(),
		Token: jwtToken,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (mi *MeInteractor) Get(authToken string) (*domain.Users, *usecase.ResultStatus) {
	db := mi.DB.Connect()

	payload, err := mi.Jwt.VerifyJwtToken(authToken)

	user, err := mi.User.FirstByID(db, payload.Audience)
	if err != nil {
		return &domain.Users{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	return user, usecase.NewResultStatus(http.StatusOK, nil)
}

func (mi *MeInteractor) Create(a *domain.SocialUserAccount) (UserResponse, *usecase.ResultStatus) {
	db := mi.DB.Begin()

	u := mi.setRegisterUser(a)

	currentTime := time.Now().Unix()
	u.CreatedAt = currentTime
	u.UpdatedAt = currentTime

	// ユーザーの新規登録
	user, err := mi.User.Create(db, u)
	if err != nil {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	// SNSアカウントとの連携
	_, err = mi.UserOauthCertification.Create(db, &domain.UserOauthCertifications{
		UserID:        user.ID,
		ServiceUserID: a.ServiceUserID,
		ServiceName:   a.ServiceName,
		CreatedAt:     currentTime,
	})
	if err != nil {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	jwtToken, payload, err := mi.Jwt.CreateToken(user.ID)
	if err != nil {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	fmt.Println(payload)
	// err = mi.Redis.Set(payload.ID, payload)
	// if err != nil {
	// 	return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

	db.Commit()
	return UserResponse{
		User:  user.BuildForGet(),
		Token: jwtToken,
	}, usecase.NewResultStatus(http.StatusAccepted, nil)
}

func (mi *MeInteractor) setRegisterUser(a *domain.SocialUserAccount) *domain.Users {
	return &domain.Users{
		DisplayName: a.DisplayName,
		ScreenName:  random.RandomScreenName(),
		Email:       a.Email,
	}
}
