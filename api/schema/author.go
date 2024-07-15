package schema

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
	Book []Book `gorm:"many2many:author_books;"`
}

type AuthorResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
