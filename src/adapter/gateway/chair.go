package gateway

import (
	"app/entity/model/chair"
	"app/entity/model/chairAuthor"
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

func (cr ChairRepository) Create(
	chairName *chair.Name,
	chairFeature *chair.Feature,
	chairYear *chair.Year,
	chairImage *chair.Image,
	chairAuthorName *chairAuthor.Name,
	chairAuthorDescription *chairAuthor.Description,
	chairAuthorBirthYear *chairAuthor.BirthYear,
	chairAuthorDiedYear *chairAuthor.DiedYear,
	chairAuthorImage *chairAuthor.Image,
) (*chair.Chair, error) {
	c := chair.Chair{
		Name:    *chairName,
		Feature: *chairFeature,
		Year:    *chairYear,
		Image:   *chairImage,
		Author: chairAuthor.ChairAuthor{
			Name:        *chairAuthorName,
			Description: *chairAuthorDescription,
			BirthYear:   *chairAuthorBirthYear,
			DiedYear:    *chairAuthorDiedYear,
			Image:       *chairAuthorImage,
		},
	}

	if err := cr.conn.Create(c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}
