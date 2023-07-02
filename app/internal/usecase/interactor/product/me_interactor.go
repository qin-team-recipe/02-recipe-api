package product

import (
	"errors"
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

func (mi *MeInteractor) LoginUser(serviceUserID string) (UserResponse, *usecase.ResultStatus) {
	db := mi.DB.Connect()

	userOauthCretificate, err := mi.UserOauthCertification.FirstByServiceUserID(db, serviceUserID)
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

func (mi *MeInteractor) Get(userID int) (*domain.Users, *usecase.ResultStatus) {
	db := mi.DB.Connect()

	// payload, err := mi.Jwt.VerifyJwtToken(authToken)

	user, err := mi.User.FirstByID(db, userID)
	if err != nil {
		return &domain.Users{}, usecase.NewResultStatus(http.StatusNotFound, err)
	}

	return user, usecase.NewResultStatus(http.StatusOK, nil)
}

func (mi *MeInteractor) Create(a *domain.SocialUserAccount) (UserResponse, *usecase.ResultStatus) {
	db := mi.DB.Begin()

	// 既存するユーザーか確認
	_, err := mi.UserOauthCertification.FirstByServiceUserID(db, a.ServiceUserID)
	if err == nil {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("既に存在するアカウントです。ログインしてください"))
	}

	// アカウント削除しているが間もなし（３０日間以内）のユーザーはどうするか

	u := mi.setRegisterUser(a)

	if u.DisplayName == "" {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("表示名が空です。"))
	}
	if u.Email == "" {
		db.Rollback()
		return UserResponse{}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("メールアドレスが空です。"))
	}

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

func (mi *MeInteractor) Save(me *domain.Users) (*domain.UsersForGet, *usecase.ResultStatus) {

	db := mi.DB.Connect()

	updatedMe, err := mi.User.Save(db, me)
	if err != nil {
		return &domain.UsersForGet{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return updatedMe.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (mi *MeInteractor) Delete(authToken string) *usecase.ResultStatus {

	db := mi.DB.Connect()

	payload, _ := mi.Jwt.VerifyJwtToken(authToken)

	user, err := mi.User.FirstByID(db, payload.Audience)
	if err != nil {
		return usecase.NewResultStatus(http.StatusNotFound, err)
	}

	if err := mi.User.Delete(db, user); err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}
	return usecase.NewResultStatus(http.StatusOK, nil)
}

func (mi *MeInteractor) setRegisterUser(a *domain.SocialUserAccount) *domain.Users {
	return &domain.Users{
		DisplayName: a.DisplayName,
		ScreenName:  random.RandomScreenName(),
		Email:       a.Email,
	}
}
