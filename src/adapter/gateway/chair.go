package gateway

import (
	"app/entity/model/chair"
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairGateway struct {
	db *gorm.DB
}

func NewChairRepository(db *gorm.DB) port.ChairRepository {
	return &chairGateway{db: db}
}

func (cr *chairGateway) Create(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *chairGateway) Delete(id *chair.ID) error {
	c := &chair.Chair{ID: *id}

	if err := cr.db.Delete(&c).Error; err != nil {
		return err
	}

	return nil
}

func (cr *chairGateway) FetchAll() ([]*chair.Chair, error) {
	var chairs []*chair.Chair
	cr.db.Joins(`ChairAuthor`).Find(&chairs)

	return chairs, nil
}

func (cr *chairGateway) FindByID(id *chair.ID) (*chair.Chair, error) {
	c := &chair.Chair{ID: *id}
	cr.db.Joins(`ChairAuthor`).Find(&c)

	return c, nil
}

func (cr *chairGateway) Update(c *chair.Chair) (*chair.Chair, error) {
	if err := cr.db.Updates(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
