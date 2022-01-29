package gateway

import (
	"app/entity/model/chair"
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairRepository struct {
	db *gorm.DB
}

func NewChairRepository(db *gorm.DB) port.ChairRepository {
	return &chairRepository{db: db}
}

func (cr *chairRepository) Create(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *chairRepository) FetchAll() ([]*chair.Chair, error) {
	var chairs []*chair.Chair
	cr.db.Joins(`ChairAuthor`).Find(&chairs)

	return chairs, nil
}

func (cr *chairRepository) FindByID(id *chair.ID) (*chair.Chair, error) {
	c := &chair.Chair{ID: *id}
	cr.db.Joins(`ChairAuthor`).Find(&c)

	return c, nil
}

func (cr *chairRepository) Update(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.db.Updates(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
