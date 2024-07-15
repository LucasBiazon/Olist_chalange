package schema

type Book struct {
	ID            uint     `gorm:"primaryKey"`
	Name          string   `gorm:"not null;size:255"`
	Edition       string   `gorm:"not null;size:255"`
	PublisherYear string   `gorm:"not null;size:255"`
	Authors       []Author `gorm:"many2many:author_books;"`
}

type BookResponse struct {
	ID            uint     `json:"id"`
	Name          string   `json:"name"`
	Edition       string   `json:"edition"`
	PublisherYear string   `json:"publisher_year"`
	Authors       []Author `json:"authors"`
}
