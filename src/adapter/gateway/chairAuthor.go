package gateway

import (
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairAuthorGateway struct {
	db *gorm.DB
}

func NewChairAuthorRepository(db *gorm.DB) port.ChairAuthorRepository {
	return &chairAuthorGateway{db: db}
}
