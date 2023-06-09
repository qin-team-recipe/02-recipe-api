package gateway

type GoogleGateway interface {
	AuthCodeURL(state string) string
	GetUserInfo(code string) (string, string, string, error)
}
