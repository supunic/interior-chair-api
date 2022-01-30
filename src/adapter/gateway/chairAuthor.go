package gateway

import (
	"app/entity/model/chairAuthor"
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairAuthorGateway struct {
	db *gorm.DB
}

func NewChairAuthorRepository(db *gorm.DB) port.ChairAuthorRepository {
	return &chairAuthorGateway{db: db}
}

func (cr *chairAuthorGateway) FetchAll() ([]*chairAuthor.ChairAuthor, error) {
	// TODO: ORMじゃないほうがドメイン周りがスッキリしそう
	var chairAuthors []*chairAuthor.ChairAuthor
	cr.db.Preload("Chairs").Find(&chairAuthors)

	return chairAuthors, nil
}
