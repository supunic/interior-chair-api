package gateway

import (
	"app/entity/model/author"
	"app/usecase/port"
	"gorm.io/gorm"
)

type authorGateway struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) port.AuthorRepository {
	return &authorGateway{db: db}
}

func (cr *authorGateway) FetchAll() ([]*author.Author, error) {
	var authors []*author.Author
	cr.db.Preload("Chairs").Find(&authors)

	return authors, nil
}
