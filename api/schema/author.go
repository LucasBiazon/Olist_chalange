package schema

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Book []Book `gorm:"many2many:author_books;"`
}

type AuthorResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
