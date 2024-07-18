package schema

import "gorm.io/gorm"

type AuthorBook struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	AuthorID string `gorm:"type:uuid"`
	BookID   string `gorm:"type:uuid"`
}
