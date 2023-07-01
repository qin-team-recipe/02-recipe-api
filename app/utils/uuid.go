package utils

import "github.com/google/uuid"

func NewUUID() string {
	uuidObj, _ := uuid.NewUUID()
	return uuidObj.String()
}

func NewRandom() string {
	uuidObj, _ := uuid.NewRandom()
	return uuidObj.String()
}
