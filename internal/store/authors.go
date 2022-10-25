package store

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/tkcrm/modules/validate"
)

type AuthorNotifications struct {
	Email bool `json:"email"`
	Sms   bool `json:"sms"`
}

func (s *CreateAuthorParams) Validate() error {
	if s.Phone != nil {
		if err := validate.Phone(*s.Phone); err != nil {
			return err
		}
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.Email, validation.Required, is.EmailFormat),
	)
}

func (s *UpdateAuthorParams) Validate() error {
	if s.ID == uuid.Nil {
		return ErrEmptyAuthorID
	}

	if s.Phone != nil {
		if err := validate.Phone(*s.Phone); err != nil {
			return err
		}
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.Email, validation.Required, is.EmailFormat),
	)
}
