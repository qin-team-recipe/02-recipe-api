package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/qin-team-recipe/02-recipe-api/config"
	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type Google struct {
	Config *oauth2.Config
}

func NewGoogle(c *config.Config) *Google {
	return newGoogle(c)
}

func newGoogle(c *config.Config) *Google {

	google := &Google{
		Config: &oauth2.Config{
			ClientID:     c.GoogleClientID,
			ClientSecret: c.GoogleSecretKey,
			Endpoint:     google.Endpoint,
			Scopes: []string{
				"openid",
				"email",
				"profile",
			},
			RedirectURL: "http://localhost:3000/auth/callback/google",
		},
	}

	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return google
}

func (g *Google) AuthCodeURL(state string) string {
	return g.Config.AuthCodeURL(state)
}

func (g *Google) GetUserInfo(code string) (string, string, string, error) {

	ctx := context.Background()

	httpClient, err := g.Config.Exchange(ctx, code)
	if err != nil {
		return "", "", "", err
	}

	client := g.Config.Client(ctx, httpClient)

	service, err := v2.New(client)
	if err != nil {
		return "", "", "", err
	}

	token, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(ctx).Do()
	if err != nil {
		return "", "", "", err
	}
	// Debug
	log.Println("token:", token)
	fmt.Printf("%+v\n", token)

	userinfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return "", "", "", err
	}

	user := &domain.Users{
		DisplayName: userinfo.Name,
		Email:       userinfo.Email,
	}
	// Debug
	log.Println("user:", user)
	fmt.Printf("%+v\n", user)

	return userinfo.Id, userinfo.Name, userinfo.Email, nil
}
