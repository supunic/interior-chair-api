package builder

import (
	"time"
)

type Chair struct {
	ID        uint      `gorm:""`
	AuthorID  uint      `gorm:""`
	Name      string    `gorm:""`
	Feature   string    `gorm:""`
	Year      int       `gorm:""`
	Image     string    `gorm:""`
	CreatedAt time.Time `gorm:""`
	UpdatedAt time.Time `gorm:""`
	Author    ChairAuthor
}

type ChairNotification interface {
	Build() *Chair
}

func NewChairRepositoryData() ChairNotification {
	return &Chair{}
}

func (c *Chair) Build() *Chair {
	return c
}
