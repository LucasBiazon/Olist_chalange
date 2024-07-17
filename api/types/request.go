package types

import (
	"fmt"

	"github.com/lucasBiazon/olist/api/schema"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// Books is a struct that represents the author entity

type CreateBookRequest struct {
	Name          string          `json:"name"`
	Edition       string          `json:"edition"`
	PublisherYear string          `json:"publisher_year"`
	AuthorId      []schema.Author `json:"author_id"`
}

func (r *CreateBookRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Edition == "" {
		return errParamIsRequired("edition", "string")
	}
	if r.PublisherYear == "" {
		return errParamIsRequired("publisher_year", "string")
	}
	if len(r.AuthorId) == 0 {
		return errParamIsRequired("author_id", "array")
	}
	return nil
}

type UpdateBookRequest struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	Edition       string          `json:"edition"`
	PublisherYear string          `json:"publisher_year"`
	AuthorId      []schema.Author `json:"author_id"`
}

func (r *UpdateBookRequest) Validate() error {
	if r.ID != "" || r.Name != "" || r.Edition != "" || r.PublisherYear != "" || len(r.AuthorId) != 0 {
		return nil
	}
	return fmt.Errorf("at least one field must be updated")
}
