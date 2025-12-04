package mocks

import (
	"errors"
)

var (
	ErrUsernameNotAllowed = errors.New("username is not allowed")
	ErrUsernameTooLong    = errors.New("username may not be longer than 20 characters")
)
