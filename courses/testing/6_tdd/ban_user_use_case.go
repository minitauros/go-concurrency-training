package tdd

import (
	"context"
)

type Status string

const (
	StatusActive Status = "active"
	StatusBanned Status = "banned"
)

type BanReason string

const (
	BanReasonRandom = "rolled_the_dice"
)

type User struct {
	ID       int64
	Username string
	Status   Status
}

func (u User) IsBanned() bool {
	return u.Status == StatusBanned
}

type UserBannedMessage struct {
	ID     int64
	Reason string
}

type UserRepo interface {
	GetUser(ctx context.Context, userID int64) (User, error)
	Ban(ctx context.Context, userID int64) error
}

type Publisher interface {
	Publish(ctx context.Context, msg []byte) error
}

type BanUserUseCase struct {
	Repo      UserRepo
	Publisher Publisher
}

func (uc *BanUserUseCase) Ban(ctx context.Context, userID int64) error {
	return nil
}
