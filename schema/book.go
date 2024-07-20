package schema

import (
	"time"
)

type Book struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Edition       string
	PublisherYear string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	Authors       []*Author `gorm:"many2many:authors_books;" json:"-"`
}
