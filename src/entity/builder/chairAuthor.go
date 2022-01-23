package builder

import (
	"time"
)

type ChairAuthor struct {
	ID          uint      `gorm:""`
	Name        string    `gorm:""`
	Description string    `gorm:""`
	BirthYear   int       `gorm:""`
	DiedYear    int       `gorm:""`
	Image       string    `gorm:""`
	CreatedAt   time.Time `gorm:""`
	UpdatedAt   time.Time `gorm:""`
}

type ChairAuthorNotification interface {
	Build() *ChairAuthor
}

func NewChairAuthorRepositoryData() ChairAuthorNotification {
	return &ChairAuthor{}
}

func (c *ChairAuthor) Build() *ChairAuthor {
	return c
}
