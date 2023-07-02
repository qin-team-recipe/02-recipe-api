package gateways

import "github.com/qin-team-recipe/02-recipe-api/pkg/token"

type JwtGateway struct {
	Jwt Jwt
}

func (j *JwtGateway) CreateToken(userID int) (string, *token.Payload, error) {
	return j.Jwt.CreateToken(userID)
}

func (j *JwtGateway) VerifyJwtToken(token string) (*token.Payload, error) {
	return j.Jwt.VerifyJwtToken(token)
}
