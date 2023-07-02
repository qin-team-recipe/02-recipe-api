package gateway

import "github.com/qin-team-recipe/02-recipe-api/pkg/token"

type JwtGateway interface {
	CreateToken(userID int) (string, *token.Payload, error)
	VerifyJwtToken(token string) (*token.Payload, error)
}
