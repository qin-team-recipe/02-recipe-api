package infrastructure

import (
	"context"
	"log"

	"github.com/qin-team-recipe/02-recipe-api/config"
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

func (g *Google) GetUserInfo(code string) (string, error) {

	cxt := context.Background()

	httpClient, err := g.Config.Exchange(cxt, code)
	if err != nil {
		return "", err
	}

	client := g.Config.Client(cxt, httpClient)

	service, err := v2.New(client)
	if err != nil {
		return "", err
	}

	userInfo, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(cxt).Do()
	if err != nil {
		return "", err
	}

	log.Println("userInfo:", userInfo)

	return "", nil
}
