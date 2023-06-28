package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/qin-team-recipe/02-recipe-api/config"
)

type JwtMaker struct {
	ApplicationName string
	TokenExpireAt   time.Duration
	SecretKey       string
}

// type UnsignedResponse struct {
// 	Message interface{} `json:"message"`
// }

func NewJwtMaker(c *config.Config) Maker {
	maker := &JwtMaker{
		ApplicationName: c.ApplicationName,
		TokenExpireAt:   c.TokenExpireAt,
		SecretKey:       c.SecretKey,
	}

	return maker
}

// Create json web token
func (j *JwtMaker) CreateToken(userID int) string {
	payload, _ := NewPayload(j.ApplicationName, userID, j.TokenExpireAt)
	// ES256 には公開鍵と秘密鍵のペアが必要で、HS256 には秘密鍵のみが必要
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Add Signature to Token
	tokenString, _ := token.SignedString([]byte(j.SecretKey))
	return tokenString
}

func (j *JwtMaker) parseToken(jwtToken string) (*Payload, error) {
	token, _ := jwt.ParseWithClaims(jwtToken, &Payload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if claims, ok := token.Claims.(*Payload); ok {
		return claims, nil
	} else {
		return claims, errors.New("invalid payload")
	}
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, ".")
	if len(jwtToken) != 3 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return header, nil
}

func (j *JwtMaker) VerifyJwtToken(token string) (*Payload, error) {
	payload, err := j.parseToken(token)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
