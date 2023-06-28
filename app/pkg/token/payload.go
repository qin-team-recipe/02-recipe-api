package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrExpiredToken = errors.New("")

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Audience  int       `json:"audience"`
	IssuedAt  int64     `json:"issued_at"`
	ExpiredAt int64     `json:"expired_at"`
}

func NewPayload(username int, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Audience:  username,
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
