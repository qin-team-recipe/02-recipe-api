package token

type Maker interface {
	CreateToken(userID int) string
	VerifyJwtToken(token string) (*Payload, error)
}
