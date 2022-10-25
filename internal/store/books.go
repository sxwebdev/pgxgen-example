package store

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

func (s *CreateBookParams) Validate() error {
	if s.AuthorID == uuid.Nil {
		return ErrEmptyAuthorID
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.Description, validation.Required),
		validation.Field(&s.ReleaseDate, validation.Required),
	)
}

func (s *UpdateBookParams) Validate() error {
	if s.ID == uuid.Nil {
		return ErrEmptyAuthorID
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.Description, validation.Required),
		validation.Field(&s.ReleaseDate, validation.Required),
	)
}
