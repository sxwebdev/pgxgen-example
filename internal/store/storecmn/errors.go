package storecmn

import "errors"

var (
	ErrEmptyID       = errors.New("empty id")
	ErrEmptyAuthorID = errors.New("empty author id")
	ErrEmptyBookID   = errors.New("empty book id")
)
