package gateways

type Google interface {
	AuthCodeURL(state string) string
	GetUserInfo(code string) (string, string, string, error)
}
