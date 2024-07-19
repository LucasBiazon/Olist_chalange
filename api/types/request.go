package types

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateBookRequest struct {
	Name          string   `json:"name" binding:"required"`
	Edition       string   `json:"edition" binding:"required"`
	PublisherYear string   `json:"publisher_year" binding:"required"`
	Authors       []string `json:"authors" binding:"required"`
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
	if len(r.Authors) == 0 {
		return errParamIsRequired("authors", "[]string")
	}
	return nil
}

type UpdateBookRequest struct {
	Name          string   `json:"name" `
	Edition       string   `json:"edition" `
	PublisherYear string   `json:"publisher_year"`
	Authors       []string `json:"authors" `
}

func (r *UpdateBookRequest) Validate() error {
	if r.Name == "" || r.Edition == "" || r.PublisherYear == "" || len(r.Authors) == 0 {
		return errParamIsRequired("name, edition, publisher_year, authors", "string, string, string, []string")
	}
	return nil
}
