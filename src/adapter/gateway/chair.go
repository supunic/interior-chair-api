package gateway

import (
	"app/entity/model"
	"app/usecase/port"
	"fmt"
	"gorm.io/gorm"
)

type ChairRepository struct {
	conn *gorm.DB
}

func NewChairRepository(conn *gorm.DB) port.ChairRepository {
	return &ChairRepository{conn: conn}
}

func (c ChairRepository) FindByID(id int) (*model.Chair, error) {
	chair := model.Chair{ID: id}

	if err := c.conn.First(chair).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("chair Not Found. id = %d", id)
		}

		return nil, fmt.Errorf("internal Server Error. %s", err.Error())
	}

	return &chair, nil
}
