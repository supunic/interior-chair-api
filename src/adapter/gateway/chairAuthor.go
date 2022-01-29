package gateway

import (
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairAuthorRepository struct {
	db *gorm.DB
}

func NewChairAuthorRepository(db *gorm.DB) port.ChairAuthorRepository {
	return &chairAuthorRepository{db: db}
}
