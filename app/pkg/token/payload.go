package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("")

type Payload struct {
	ID        string `json:"id"`
	Audience  int    `json:"audience"`
	Issuer    string `json:"issuer"`
	Subject   string `json:"subject"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiredAt int64  `json:"expired_at"`
}

func NewPayload(applicationName string, username int, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	fmt.Println(duration)

	payload := &Payload{
		ID:        tokenID.String(),
		Audience:  username,
		Issuer:    applicationName,
		Subject:   "http://localhost:3000",
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(duration).Unix(),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(time.Unix(payload.ExpiredAt, 0)) {
		return ErrExpiredToken
	}
	return nil
}
