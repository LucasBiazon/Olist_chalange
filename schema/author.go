package schema

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID        string `gorm:"type:uuid;primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Books     []*Book        `gorm:"many2many:authors_books;"`
}
