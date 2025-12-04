package mocks

import (
	"context"
	"unicode/utf8"
)

type User struct {
	ID       int64
	Username string
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) (id int64, err error)
}

type CreateUserUseCase struct {
	// Repo is the repository used to create users. This field would usually not be exported,
	// but now _is_ exported so that the implementation of the solution can have access to it.
	Repo UserRepository
}

func (uc *CreateUserUseCase) CreateUser(ctx context.Context, username string) (id int64, err error) {
	if username == "disallowed" {
		return 0, ErrUsernameNotAllowed
	} else if utf8.RuneCountInString(username) > 20 {
		return 0, ErrUsernameTooLong
	}

	id, err = uc.Repo.CreateUser(ctx, User{
		Username: username,
	})

	return
}
