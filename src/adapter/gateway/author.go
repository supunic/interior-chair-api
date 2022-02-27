package gateway

import (
	"app/usecase/data"
	"app/usecase/port"
	"gorm.io/gorm"
)

type authorGateway struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) port.AuthorRepository {
	return &authorGateway{db: db}
}

func (cr *authorGateway) FetchAll() ([]*data.AuthorRepositoryData, error) {
	var as []*data.AuthorRepositoryData

	if err := cr.db.Find(&as).Error; err != nil {
		return nil, err
	}

	return as, nil
}
