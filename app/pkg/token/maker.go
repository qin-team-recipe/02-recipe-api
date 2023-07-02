package token

type Maker interface {
	CreateToken(userID int) (string, *Payload, error)
	VerifyJwtToken(token string) (*Payload, error)
}
