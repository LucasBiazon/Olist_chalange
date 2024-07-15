package schema

import "gorm.io/gorm"

type AuthorBook struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	AuthorID uint
	BookID   uint
}

type AuthorBookResponse struct {
	ID       uint `json:"id"`
	AuthorID uint `json:"author_id"`
	BookID   uint `json:"book_id"`
}
