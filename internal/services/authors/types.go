package authors

import (
	"github.com/google/uuid"
	"github.com/sxwebdev/pgxgen-example/internal/models"
)

type CreateParams struct {
	FirstName     string                     `json:"first_name"`
	LastName      string                     `json:"last_name"`
	Email         string                     `json:"email"`
	Phone         *string                    `json:"phone"`
	IsActive      bool                       `json:"is_active"`
	Notifications models.AuthorNotifications `json:"notifications"`
}

type UpdateParams struct {
	FirstName     string                     `json:"first_name"`
	LastName      string                     `json:"last_name"`
	Email         string                     `json:"email"`
	Phone         *string                    `json:"phone"`
	IsActive      bool                       `json:"is_active"`
	Notifications models.AuthorNotifications `json:"notifications"`
	ID            uuid.UUID                  `json:"id"`
}
