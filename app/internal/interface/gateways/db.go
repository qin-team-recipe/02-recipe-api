package gateways

import "gorm.io/gorm"

type DB interface {
	Connect() *gorm.DB
	Begin() *gorm.DB
}
