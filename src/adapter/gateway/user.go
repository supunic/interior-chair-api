package gateway

import (
	"app/usecase/port"
	"gorm.io/gorm"
)

type userGateway struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &userGateway{db: db}
}
