package gateway

import (
	"app/usecase/data"
	"app/usecase/port"
	"gorm.io/gorm"
)

type chairGateway struct {
	db *gorm.DB
}

func NewChairRepository(db *gorm.DB) port.ChairRepository {
	return &chairGateway{db: db}
}

func (cr *chairGateway) Create(c *data.ChairRepositoryData) (*data.ChairRepositoryData, error) {
	if err := cr.db.Create(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *chairGateway) Delete(c *data.ChairRepositoryData) error {
	if err := cr.db.Delete(&c).Error; err != nil {
		return err
	}

	return nil
}

func (cr *chairGateway) FetchAll() ([]*data.ChairRepositoryData, error) {
	var cs []*data.ChairRepositoryData

	if err := cr.db.Joins(`Author`).Find(&cs).Error; err != nil {
		return nil, err
	}

	return cs, nil
}

func (cr *chairGateway) FindByID(c *data.ChairRepositoryData) (*data.ChairRepositoryData, error) {
	if err := cr.db.Joins(`Author`).Find(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}

func (cr *chairGateway) Update(c *data.ChairRepositoryData) (*data.ChairRepositoryData, error) {
	if err := cr.db.Updates(&c).Error; err != nil {
		return nil, err
	}

	return c, nil
}
