package gateways

type GoogleGateway struct {
	Google Google
}

func (g *GoogleGateway) AuthCodeURL(state string) string {
	return g.Google.AuthCodeURL(state)
}

func (g *GoogleGateway) GetUserInfo(code string) (string, error) {
	return g.Google.GetUserInfo(code)
}
