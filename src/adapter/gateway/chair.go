package gateway

import (
	"app/entity/model/chair"
	"app/usecase/port"
	"gorm.io/gorm"
)

type ChairRepository struct {
	db *gorm.DB
}

func NewChairRepository(db *gorm.DB) port.ChairRepository {
	return &ChairRepository{db: db}
}

func (cr *ChairRepository) Create(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *ChairRepository) FindByID(id *chair.ID) (*chair.Chair, error) {
	c := &chair.Chair{ID: *id}
	cr.db.Joins(`ChairAuthor`).Find(&c)

	return c, nil
}
