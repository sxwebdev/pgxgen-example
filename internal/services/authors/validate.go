package authors

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/store/storecmn"
	"github.com/tkcrm/modules/pkg/validate"
)

func (s *CreateParams) Validate() error {
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

func (s *UpdateParams) Validate() error {
	if s.ID == uuid.Nil {
		return storecmn.ErrEmptyID
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
