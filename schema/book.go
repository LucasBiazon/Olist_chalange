package schema

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            string `gorm:"primaryKey"`
	Name          string
	Edition       string
	PublisherYear string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Authors       []*Author `gorm:"many2many:authors_books;" json:"Authors"`
}
