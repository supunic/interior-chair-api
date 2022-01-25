package gateway

import (
	"app/entity/model/chair"
	"app/usecase/port"
	"gorm.io/gorm"
)

type ChairRepository struct {
	conn *gorm.DB
}

func NewChairRepository(conn *gorm.DB) port.ChairRepository {
	return &ChairRepository{conn: conn}
}

func (cr *ChairRepository) Create(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.conn.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *ChairRepository) FindByID(id *chair.ID) (*chair.Chair, error) {
	c := &chair.Chair{ID: *id}
	cr.conn.Joins(`ChairAuthor`).Find(&c)

	return c, nil
}
