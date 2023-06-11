package gateway

import "gorm.io/gorm"

type DBRepository interface {
	Connect() *gorm.DB
	Begin() *gorm.DB
}
