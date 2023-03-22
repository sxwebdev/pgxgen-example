package books

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/store/storecmn"
)

func (s *CreateParams) Validate() error {
	if s.AuthorID == uuid.Nil {
		return storecmn.ErrEmptyAuthorID
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.Description, validation.Required),
		validation.Field(&s.ReleaseDate, validation.Required),
	)
}

func (s *UpdateParams) Validate() error {
	if s.ID == uuid.Nil {
		return storecmn.ErrEmptyAuthorID
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.Description, validation.Required),
		validation.Field(&s.ReleaseDate, validation.Required),
	)
}
