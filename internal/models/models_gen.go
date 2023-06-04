// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package models

import (
	"database/sql/driver"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BookType string

const (
	BookTypeAdventure BookType = "adventure"
	BookTypeNovel     BookType = "novel"
	BookTypeDetective BookType = "detective"
)

func (e *BookType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = BookType(s)
	case string:
		*e = BookType(s)
	default:
		return fmt.Errorf("unsupported scan type for BookType: %T", src)
	}
	return nil
}

type NullBookType struct {
	BookType BookType
	Valid    bool // Valid is true if BookType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullBookType) Scan(value interface{}) error {
	if value == nil {
		ns.BookType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.BookType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullBookType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.BookType), nil
}

func (e BookType) Valid() bool {
	switch e {
	case BookTypeAdventure,
		BookTypeNovel,
		BookTypeDetective:
		return true
	}
	return false
}

func AllBookTypeValues() []BookType {
	return []BookType{
		BookTypeAdventure,
		BookTypeNovel,
		BookTypeDetective,
	}
}

type Author struct {
	ID            uuid.UUID           `db:"id" json:"id"`
	FirstName     string              `db:"first_name" json:"first_name"`
	LastName      string              `db:"last_name" json:"last_name"`
	Email         string              `db:"email" json:"email"`
	Phone         pgtype.Text         `db:"phone" json:"phone"`
	IsActive      bool                `db:"is_active" json:"is_active"`
	Notifications AuthorNotifications `db:"notifications" json:"notifications"`
	CreatedAt     pgtype.Timestamp    `db:"created_at" json:"created_at"`
	UpdatedAt     pgtype.Timestamp    `db:"updated_at" json:"updated_at"`
}

type Book struct {
	ID          uuid.UUID        `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Description string           `db:"description" json:"description"`
	Genre       BookType         `db:"genre" json:"genre"`
	ReleaseDate pgtype.Timestamp `db:"release_date" json:"release_date"`
	AuthorID    uuid.UUID        `db:"author_id" json:"author_id"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"created_at"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updated_at"`
}
