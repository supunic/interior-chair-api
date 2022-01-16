package gateway

import (
	"app/entity/model/chair"
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

func (cr ChairRepository) FindByID(id *chair.ID) (*chair.Chair, error) {
	c := chair.Chair{ID: *id}

	if err := cr.conn.First(c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("chair Not Found. id = %d", id)
		}

		return nil, fmt.Errorf("internal Server Error. %s", err.Error())
	}

	return &c, nil
}
