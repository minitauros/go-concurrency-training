package tdd_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	tdd "github.com/minitauros/go-concurrency-training/courses/testing/6_tdd"
	"github.com/minitauros/go-concurrency-training/courses/testing/6_tdd/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_BanUserUseCase_Ban(t *testing.T) {
	t.Run("if an error happens while fetching the user from the db, returns error", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepo(t)
		publisher := mocks.NewMockPublisher(t)
		uc := &tdd.BanUserUseCase{
			Repo:      userRepo,
			Publisher: publisher,
		}

		ctx := context.Background()
		userID := int64(123)
		expectedErr := errors.New("err")

		userRepo.EXPECT().GetUser(ctx, userID).Return(tdd.User{}, expectedErr).Once()

		err := uc.Ban(ctx, userID)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("if the user is already banned, returns nil", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepo(t)
		publisher := mocks.NewMockPublisher(t)
		uc := &tdd.BanUserUseCase{
			Repo:      userRepo,
			Publisher: publisher,
		}

		ctx := context.Background()
		user := tdd.User{
			ID:       123,
			Username: "foo",
			Status:   tdd.StatusBanned,
		}

		userRepo.EXPECT().GetUser(ctx, user.ID).Return(user, nil).Once()

		err := uc.Ban(ctx, user.ID)
		assert.Nil(t, err)
	})

	t.Run("if the user is not already banned, but banning fails, returns error", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepo(t)
		publisher := mocks.NewMockPublisher(t)
		uc := &tdd.BanUserUseCase{
			Repo:      userRepo,
			Publisher: publisher,
		}

		ctx := context.Background()
		user := tdd.User{
			ID:       123,
			Username: "foo",
			Status:   tdd.StatusActive,
		}
		expectedErr := errors.New("err")

		userRepo.EXPECT().GetUser(ctx, user.ID).Return(user, nil).Once()
		userRepo.EXPECT().Ban(ctx, user.ID).Return(expectedErr).Once()

		err := uc.Ban(ctx, user.ID)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("if banning succeeds, but publishing the 'user banned' message fails, returns error", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepo(t)
		publisher := mocks.NewMockPublisher(t)
		uc := &tdd.BanUserUseCase{
			Repo:      userRepo,
			Publisher: publisher,
		}

		ctx := context.Background()
		user := tdd.User{
			ID:       123,
			Username: "foo",
			Status:   tdd.StatusActive,
		}
		expectedErr := errors.New("err")
		userBannedMessage := tdd.UserBannedMessage{
			ID:     user.ID,
			Reason: tdd.BanReasonRandom,
		}
		userBannedMessageBytes, err := json.Marshal(userBannedMessage)
		if err != nil {
			panic(err)
		}

		userRepo.EXPECT().GetUser(ctx, user.ID).Return(user, nil).Once()
		userRepo.EXPECT().Ban(ctx, user.ID).Return(nil).Once()
		publisher.EXPECT().Publish(ctx, userBannedMessageBytes).Return(expectedErr).Once()

		err = uc.Ban(ctx, user.ID)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("if publishing the 'user banned' message succeeds, returns nil", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepo(t)
		publisher := mocks.NewMockPublisher(t)
		uc := &tdd.BanUserUseCase{
			Repo:      userRepo,
			Publisher: publisher,
		}

		ctx := context.Background()
		user := tdd.User{
			ID:       123,
			Username: "foo",
			Status:   tdd.StatusActive,
		}
		userBannedMessage := tdd.UserBannedMessage{
			ID:     user.ID,
			Reason: tdd.BanReasonRandom,
		}
		userBannedMessageBytes, err := json.Marshal(userBannedMessage)
		if err != nil {
			panic(err)
		}

		userRepo.EXPECT().GetUser(ctx, user.ID).Return(user, nil).Once()
		userRepo.EXPECT().Ban(ctx, user.ID).Return(nil).Once()
		publisher.EXPECT().Publish(ctx, userBannedMessageBytes).Return(nil).Once()

		err = uc.Ban(ctx, user.ID)
		assert.Nil(t, err)
	})
}
