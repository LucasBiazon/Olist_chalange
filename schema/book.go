package schema

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            string `gorm:"type:uuid;primaryKey"`
	Name          string `gorm:"not null;size:255"`
	Edition       string `gorm:"not null;size:255"`
	PublisherYear string `gorm:"not null;size:255"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Authors       []*Author      `gorm:"many2many:authors_books;"`
}
