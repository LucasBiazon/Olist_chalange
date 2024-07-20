package schema

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Books     []*Book `gorm:"many2many:authors_books;" json:"-"`
}
